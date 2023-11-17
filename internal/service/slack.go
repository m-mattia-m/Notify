package service

import (
	"github.com/slack-go/slack"
)

type SlackService interface {
	SendMessage(to, message, token string) error
}

type SlackClient struct{}

func NewSlackClient() SlackService {
	return &SlackClient{}
}

func (sv *SlackClient) SendMessage(to, message, token string) error {
	api := slack.New(token)
	_, _, err := api.PostMessage(to, slack.MsgOptionText(message, false))
	return err
}
