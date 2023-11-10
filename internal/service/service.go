package service

type Client struct {
	mailgun MailgunService
	slack   SlackService
	db      DbService
}

type Service interface {
	IfHostOrIpVerified(clientIP string, clientHost string) (bool, error)
}

func NewClient() Service {
	return &Client{
		mailgun: NewMailgunClient(),
		slack:   NewSlackClient(),
		db:      NewDbClient(),
	}
}
