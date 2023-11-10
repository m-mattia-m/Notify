package service

import log "github.com/sirupsen/logrus"

type SlackService interface {
	SendMessage() error
}

type SlackClient struct{}

func NewSlackClient() SlackService {
	return &SlackClient{}
}

func (sv *SlackClient) SendMessage() error {
	log.Infof("Sending message via Slack")
	return nil
}
