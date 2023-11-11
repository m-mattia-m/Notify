package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"message-proxy/internal/dao"
	"message-proxy/internal/model"
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
}

type DbClient struct {
	dao dao.Dao
}

func NewDbClient() DbService {
	var uri = fmt.Sprintf("mongodb://%s:%s", viper.GetString("db.mongo.host"), viper.GetString("db.mongo.port"))
	if viper.GetString("app.env") == "PROD" {
		uri = fmt.Sprintf("mongodb+srv://%s/?tls=true", viper.GetString("db.mongo.host"))
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	credentials := options.Credential{
		Username: viper.GetString("db.mongo.user"),
		Password: viper.GetString("db.mongo.password"),
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
		dao: dao.New(client, viper.GetString("db.mongo.name")),
	}
}
