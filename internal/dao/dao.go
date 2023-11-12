package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"notify/internal/model"
)

type Dao interface {
	GetConnection() error

	IfHostVerified(clientHost string) (bool, error)
	IfHostInThisProjectAlreadyExist(host model.Host) (bool, error)
	CreateHost(host model.Host) (*model.Host, error)
	GetHost(hostFilter model.Host) (*model.Host, error)
	ListHosts(hostFilter model.Host) ([]*model.Host, error)
	UpdateHost(host model.Host) (*model.Host, error)
	DeleteHost(hostFilter model.Host) error

	IfProjectWithThisNameAlreadyExist(project model.Project) (bool, error)
	CreateProject(project model.Project) (*model.Project, error)
	GetProject(projectFilter model.Project) (*model.Project, error)
	ListProjects(projectFilter model.Project) ([]*model.Project, error)
	UpdateProject(project model.Project) (*model.Project, error)
	DeleteProject(projectFilter model.Project) error

	CreateSlackCredential(credentials model.SlackCredentials) error
	IsSlackCredentialsAlreadySet(credentials model.SlackCredentials) (bool, error)
	UpdateSlackCredential(credentials model.SlackCredentials) error
	DeleteSlackCredential(credentials model.SlackCredentials) error

	CreateMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error)
	GetMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error)
	IsMailgunCredentialsAlreadySet(credentials model.MailgunCredentials) (bool, error)
	UpdateMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error)
	DeleteMailgunCredential(credentials model.MailgunCredentials) error
}

type DaoClient struct {
	engine *mongo.Client
	dbName string
}

func New(engine *mongo.Client, dbName string) Dao {
	return &DaoClient{
		engine: engine,
		dbName: dbName,
	}
}

func (dc *DaoClient) GetConnection() error {
	err := dc.engine.Ping(context.Background(), &readpref.ReadPref{})
	return err
}
