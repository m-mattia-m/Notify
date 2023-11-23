package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"notify/internal/model"
)

func (c *Client) GetMailgunCredentials(projectId string) (*model.MailgunCredentialsResponse, error) {

	credentials, err := c.db.GetMailgunCredential(model.MailgunCredentials{ProjectId: projectId})
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return credentials, nil
}

func (c *Client) CreateMailgunCredentials(projectId string, credentialsRequest model.MailgunCredentialsRequest) (*model.MailgunCredentialsResponse, error) {
	err := validateMailgunCredentialsRequest(credentialsRequest)
	if err != nil {
		return nil, err
	}

	credentials := model.MailgunCredentials{
		ProjectId:   projectId,
		Domain:      credentialsRequest.Domain,
		ApiKey:      credentialsRequest.ApiKey,
		SenderEmail: credentialsRequest.SenderEmail,
		SenderName:  credentialsRequest.SenderName,
	}

	createdCredentials, err := c.db.CreateMailgunCredential(credentials)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return createdCredentials, nil
}

func (c *Client) UpdateMailgunCredentials(projectId string, credentialsRequest model.MailgunCredentialsRequest) (*model.MailgunCredentialsResponse, error) {
	err := validateMailgunCredentialsRequest(credentialsRequest)
	if err != nil {
		return nil, err
	}

	credentials := model.MailgunCredentials{
		ProjectId:   projectId,
		Domain:      credentialsRequest.Domain,
		ApiKey:      credentialsRequest.ApiKey,
		SenderEmail: credentialsRequest.SenderEmail,
		SenderName:  credentialsRequest.SenderName,
	}

	updatedCredentials, err := c.db.UpdateMailgunCredential(credentials)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return updatedCredentials, nil

}

func (c *Client) DeleteMailgunCredentials(projectId string) (*model.SuccessMessage, error) {
	credentials := model.MailgunCredentials{
		ProjectId: projectId,
	}

	if err := c.db.DeleteMailgunCredential(credentials); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return &model.SuccessMessage{
		Message: "successfully deleted",
	}, nil

}

func (c *Client) IsMailgunCredentialsAlreadySet(projectId string) (bool, error) {
	credentials := model.MailgunCredentials{
		ProjectId: projectId,
	}

	exist, err := c.db.IsMailgunCredentialsAlreadySet(credentials)
	if err != nil {
		log.Error(err)
		return false, fmt.Errorf("")
	}

	if exist {
		return true, nil
	}

	return false, nil
}

func validateMailgunCredentialsRequest(mailgunCredentialsRequest model.MailgunCredentialsRequest) error {
	if mailgunCredentialsRequest.Domain == "" {
		return fmt.Errorf("domain is a required attribute")
	}
	if mailgunCredentialsRequest.ApiKey == "" {
		return fmt.Errorf("api-key is a required attribute")
	}
	if mailgunCredentialsRequest.ApiBase == "" {
		return fmt.Errorf("api-base is a required attribute; only 'us' or 'eu' are valid")
	}
	if mailgunCredentialsRequest.SenderEmail == "" {
		return fmt.Errorf("sender-email is a required attribute")
	}
	if mailgunCredentialsRequest.SenderName == "" {
		return fmt.Errorf("sender-name is a required attribute")
	}
	if mailgunCredentialsRequest.ReplyToEmail == "" {
		return fmt.Errorf("reply-to-email is a required attribute")
	}
	return nil
}
