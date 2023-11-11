package service

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message-proxy/internal/model"
)

func (c *Client) IfHostOrIpVerified(clientIP, clientHost string) (bool, error) {
	return c.db.IfHostOrIpVerified(clientIP, clientHost)
}

func (c *Client) CreateHost(hostRequest model.HostRequest, projectId string) (*model.Host, error) {
	host := model.Host{
		ProjectId:   projectId,
		Host:        hostRequest.Host,
		Stage:       hostRequest.Stage,
		VerifyToken: uuid.NewString(),
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

	return c.db.ListHosts(hostFilter)
}

func (c *Client) VerifyHost(hostId, projectId string) (*model.Host, error) {
	hostObjectId, err := primitive.ObjectIDFromHex(hostId)
	if err != nil {
		return nil, fmt.Errorf("invalid hostId")
	}

	host := model.Host{
		Id:        hostObjectId,
		ProjectId: projectId,
	}
	return c.db.UpdateHost(host)
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
