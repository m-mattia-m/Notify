package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Dao interface {
	GetConnection() error

	IfHostVerified(clientIP, clientHost string) (bool, error)
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
	err := dc.engine.Ping(context.TODO(), &readpref.ReadPref{})
	return err
}
