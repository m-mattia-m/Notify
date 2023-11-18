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

	CreateSlackCredentials(projectId string, credentialsRequest model.SlackCredentialsRequest) (*model.SuccessMessage, error)
	UpdateSlackCredentials(projectId string, credentialsRequest model.SlackCredentialsRequest) (*model.SuccessMessage, error)
	DeleteSlackCredentials(projectId string) (*model.SuccessMessage, error)
	IsSlackCredentialsAlreadySet(projectId string) (bool, error)

	GetMailgunCredentials(projectId string) (*model.MailgunCredentialsResponse, error)
	CreateMailgunCredentials(projectId string, credentialsRequest model.MailgunCredentialsRequest) (*model.MailgunCredentialsResponse, error)
	UpdateMailgunCredentials(projectId string, credentialsRequest model.MailgunCredentialsRequest) (*model.MailgunCredentialsResponse, error)
	DeleteMailgunCredentials(projectId string) (*model.SuccessMessage, error)
	IsMailgunCredentialsAlreadySet(projectId string) (bool, error)

	CreateFlow(projectId string, flowRequest model.FlowRequest) (*model.Flow, error)
	GetFlow(flowId string, projectId string) (*model.Flow, error)
	ListFlows(projectId string) ([]*model.Flow, error)
	UpdateFlow(flowId, projectId string, flowRequest model.FlowRequest) (*model.Flow, error)
	DeleteFlow(flowId string, projectId string) (*model.SuccessMessage, error)
	proveMessageTypes(messageType string) bool

	SendNotification(host string, notification model.Notification) (*model.SuccessMessage, error)
	sendSlackNotification(flow model.Flow, notification model.Notification, credential model.SlackCredentials) error
	sendMailgunNotification(flow model.Flow, notification model.Notification, credential model.MailgunCredentials) error

	GetActivity(activityId, projectId string) (*model.Activity, error)
	ListActivities(projectId string) ([]*model.Activity, error)
	logNotificationActivity(activity model.Activity) error
}

func NewClient() Service {
	return &Client{
		mailgun: NewMailgunClient(),
		slack:   NewSlackClient(),
		db:      NewDbClient(),
	}
}
