package service

import log "github.com/sirupsen/logrus"

type MailgunService interface {
	SendMail() error
}

type MailgunClient struct{}

func NewMailgunClient() MailgunService {
	return &MailgunClient{}
}

func (mc *MailgunClient) SendMail() error {
	log.Infof("Sending mail via Mailgun")
	return nil
}
