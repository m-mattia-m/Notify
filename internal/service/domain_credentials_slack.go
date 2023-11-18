package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"notify/internal/model"
)

func (c *Client) CreateSlackCredentials(projectId string, credentialsRequest model.SlackCredentialsRequest) (*model.SuccessMessage, error) {
	credentials := model.SlackCredentials{
		ProjectId:         projectId,
		BotUserOAuthToken: credentialsRequest.BotUserOAuthToken,
	}

	if err := c.db.CreateSlackCredential(credentials); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return &model.SuccessMessage{
		Message: "successfully created",
	}, nil
}

func (c *Client) UpdateSlackCredentials(projectId string, credentialsRequest model.SlackCredentialsRequest) (*model.SuccessMessage, error) {
	credentials := model.SlackCredentials{
		ProjectId:         projectId,
		BotUserOAuthToken: credentialsRequest.BotUserOAuthToken,
	}

	if err := c.db.UpdateSlackCredential(credentials); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return &model.SuccessMessage{
		Message: "successfully updated",
	}, nil

}

func (c *Client) DeleteSlackCredentials(projectId string) (*model.SuccessMessage, error) {
	credentials := model.SlackCredentials{
		ProjectId: projectId,
	}

	if err := c.db.DeleteSlackCredential(credentials); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return &model.SuccessMessage{
		Message: "successfully deleted",
	}, nil

}

func (c *Client) IsSlackCredentialsAlreadySet(projectId string) (bool, error) {
	credentials := model.SlackCredentials{
		ProjectId: projectId,
	}

	exist, err := c.db.IsSlackCredentialsAlreadySet(credentials)
	if err != nil {
		log.Error(err)
		return false, fmt.Errorf("")
	}

	if exist {
		return true, nil
	}

	return false, nil
}
