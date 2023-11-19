package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"notify/internal/dao"
	"notify/internal/model"
	"os"
	"strings"
)

type DbService interface {
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

	GetCredential(credentials model.Credential) (*model.Credential, error)
	ListCredential(credentialFilter model.Credential) ([]*model.Credential, error)

	CreateSlackCredential(credentials model.SlackCredentials) error
	IsSlackCredentialsAlreadySet(credentials model.SlackCredentials) (bool, error)
	UpdateSlackCredential(credentials model.SlackCredentials) error
	DeleteSlackCredential(credentials model.SlackCredentials) error
	GetSlackRevealedCredential(credentialsFilter model.SlackCredentials) (*model.SlackCredentials, error)

	CreateMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error)
	GetMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error)
	IsMailgunCredentialsAlreadySet(credentials model.MailgunCredentials) (bool, error)
	UpdateMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error)
	DeleteMailgunCredential(credentials model.MailgunCredentials) error
	GetMailgunRevealedCredential(credentials model.MailgunCredentials) (*model.MailgunCredentials, error)

	IfFlowInThisProjectAlreadyExist(flowFilter model.Flow) (bool, error)
	CreateFlow(flow model.Flow) (*model.Flow, error)
	GetFlow(flowFilter model.Flow) (*model.Flow, error)
	ListFlows(flowFilter model.Flow) ([]*model.Flow, error)
	UpdateFlow(flow model.Flow) (*model.Flow, error)
	DeleteFlow(flowFilter model.Flow) error

	CreateActivity(activity model.Activity) error
	GetActivity(activityFilter model.Activity) (*model.Activity, error)
	ListActivities(activityFilter model.Activity) ([]*model.Activity, error)
}

type DbClient struct {
	dao dao.Dao
}

func NewDbClient() DbService {
	mongoHost, found := os.LookupEnv("MONGO_HOST")
	if !found {
		log.Fatal("Error starting mongoDB-client: env 'MONGO_HOST' not found")
	}

	mongoPort, found := os.LookupEnv("MONGO_PORT")
	if !found {
		log.Info("env 'MONGO_PORT' not found, if this is not needed, you can ignore this info")
	}

	mongoUsername, found := os.LookupEnv("MONGO_USERNAME")
	if !found {
		log.Fatal("Error starting mongoDB-client: env 'MONGO_USERNAME' not found")
	}

	mongoPassword, found := os.LookupEnv("MONGO_PASSWORD")
	if !found {
		log.Fatal("Error starting mongoDB-client: env 'MONGO_PASSWORD' not found")
	}

	mongoDatabaseName, found := os.LookupEnv("MONGO_DATABASE_NAME")
	if !found {
		log.Fatal("Error starting mongoDB-client: env 'MONGO_DATABASE_NAME' not found")
	}

	mongoTlsActive, found := os.LookupEnv("MONGO_TLS_ACTIVE")
	if !found {
		log.Info("env 'MONGO_DATABASE_NAME' not found, if this is not needed, you can ignore this info")
	}

	var uri = fmt.Sprintf("mongodb://%s", mongoHost)
	if strings.Contains(mongoHost, "+srv") {
		uri = fmt.Sprintf("mongodb+srv://%s", mongoHost)
	}
	if mongoPort != "" {
		uri = fmt.Sprintf("%s:%s", uri, mongoPort)
	}
	if mongoTlsActive == "true" {
		uri = fmt.Sprintf("%s/?tls=true", uri)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	credentials := options.Credential{
		Username: mongoUsername,
		Password: mongoPassword,
	}

	if viper.GetString("app.env") != "PROD" {
		credentials.AuthMechanism = "SCRAM-SHA-256"
	}

	opts := options.Client().
		ApplyURI(uri).
		SetAuth(credentials).
		SetServerAPIOptions(serverAPI)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect with the mongoDB")
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("failed to ping the mongoDB")
	}

	return &DbClient{
		dao: dao.New(client, mongoDatabaseName),
	}
}
