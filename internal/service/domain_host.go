package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message-proxy/internal/model"
	"slices"
	"strings"
)

const GOOGLE_DNS_SERVER = "8.8.8.8:53"

func (c *Client) IfHostVerified(clientHost string) (bool, error) {
	return c.db.IfHostVerified(clientHost)
}

func (c *Client) CreateHost(hostRequest model.HostRequest, projectId string) (*model.Host, error) {
	host := model.Host{
		ProjectId:   projectId,
		Host:        hostRequest.Host,
		Stage:       hostRequest.Stage,
		VerifyToken: fmt.Sprintf("notify-verification::%s", uuid.NewString()),
	}

	alreadyExist, err := c.db.IfHostInThisProjectAlreadyExist(host)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	if alreadyExist {
		return nil, fmt.Errorf("host already exist in this project")
	}

	createdHost, err := c.db.CreateHost(host)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return createdHost, nil
}

func (c *Client) GetHost(hostId, projectId string) (*model.Host, error) {
	hostObjectId, err := primitive.ObjectIDFromHex(hostId)
	if err != nil {
		return nil, fmt.Errorf("invalid hostId")
	}

	hostFilter := model.Host{
		Id:        hostObjectId,
		ProjectId: projectId,
	}

	host, err := c.db.GetHost(hostFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return host, nil
}

func (c *Client) ListHosts(projectId string) ([]*model.Host, error) {
	hostFilter := model.Host{
		ProjectId: projectId,
	}
	hosts, err := c.db.ListHosts(hostFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return hosts, err
}

func (c *Client) VerifyHost(hostId, projectId string) (*model.Host, error) {
	hostObjectId, err := primitive.ObjectIDFromHex(hostId)
	if err != nil {
		return nil, fmt.Errorf("invalid hostId")
	}

	hostFilter := model.Host{
		Id:        hostObjectId,
		ProjectId: projectId,
	}
	host, err := c.db.GetHost(hostFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	if host.Verified {
		return nil, fmt.Errorf("already verified")
	}

	if strings.ToLower(host.Stage) == "local" && strings.HasPrefix(host.Host, "localhost:") {
		toUpdatedHost := model.Host{
			Id:        hostObjectId,
			ProjectId: projectId,
			Verified:  true,
		}
		updatedHost, err := c.db.UpdateHost(toUpdatedHost)
		if err != nil {
			log.Error(err)
			return nil, fmt.Errorf("")
		}

		return updatedHost, nil
	}

	verified, err := c.verifyHost(host.Host, host.VerifyToken)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	if !verified {
		return nil, fmt.Errorf("failed to verify the host, proove your DNS config")
	}

	toUpdatedHost := model.Host{
		Id:        hostObjectId,
		ProjectId: projectId,
		Verified:  verified,
	}
	updatedHost, err := c.db.UpdateHost(toUpdatedHost)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return updatedHost, nil
}

func (c *Client) DeleteHost(hostId, projectId string) (*model.SuccessMessage, error) {

	hostObjectId, err := primitive.ObjectIDFromHex(hostId)
	if err != nil {
		return nil, fmt.Errorf("invalid hostId")
	}

	hostFilter := model.Host{
		Id:        hostObjectId,
		ProjectId: projectId,
	}
	hostResponse, err := c.db.GetHost(hostFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	err = c.db.DeleteHost(hostFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return &model.SuccessMessage{
		Message: fmt.Sprintf("%s successfully deleted", hostResponse.Host),
	}, nil
}

func (c *Client) verifyHost(host, verificationToken string) (bool, error) {
	dnsServer := GOOGLE_DNS_SERVER
	if customDns := viper.GetString("domain.dns.verifyDns"); customDns != "" {
		dnsServer = customDns
	}

	txtRecords, err := c.queryTXTVerificationRecord(host, dnsServer)
	if err != nil {
		return false, err
	}

	if slices.Contains(txtRecords, verificationToken) {
		return true, nil
	}
	return false, nil
}

func (c *Client) queryTXTVerificationRecord(host, dnsServer string) ([]string, error) {
	dnsClient := new(dns.Client)
	dnsMessage := new(dns.Msg)
	dnsMessage.SetQuestion(dns.Fqdn(host), dns.TypeTXT)

	dnsResponse, _, err := dnsClient.Exchange(dnsMessage, dnsServer)
	if err != nil {
		return nil, err
	}

	var txtRecords []string
	for _, ans := range dnsResponse.Answer {
		if txt, ok := ans.(*dns.TXT); ok {
			txtRecords = append(txtRecords, txt.Txt...)
		}
	}

	return txtRecords, nil
}
