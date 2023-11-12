package service

import "notify/internal/model"

type Client struct {
	mailgun MailgunService
	slack   SlackService
	db      DbService
}

type Service interface {
	IfHostVerified(clientHost string) (bool, error)
	CreateHost(hostRequest model.HostRequest, projectId string) (*model.Host, error)
	GetHost(hostId, projectId string) (*model.Host, error)
	ListHosts(projectId string) ([]*model.Host, error)
	VerifyHost(hostId, projectId string) (*model.Host, error)
	DeleteHost(hostId, projectId string) (*model.SuccessMessage, error)

	CreateProject(projectRequest model.ProjectRequest, userId string) (*model.Project, error)
	GetProject(projectId, userId string) (*model.Project, error)
	ListProjects(userId string) ([]*model.Project, error)
	UpdateProject(projectId string, projectRequest model.ProjectRequest, userId string) (*model.Project, error)
	DeleteProject(projectId, userId string) (*model.SuccessMessage, error)

	verifyHost(host, verificationToken string) (bool, error)
	queryTXTVerificationRecord(host, dnsServer string) ([]string, error)
}

func NewClient() Service {
	return &Client{
		mailgun: NewMailgunClient(),
		slack:   NewSlackClient(),
		db:      NewDbClient(),
	}
}
