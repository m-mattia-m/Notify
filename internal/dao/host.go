package dao

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"notify/internal/model"
	"time"
)

func (dc *DaoClient) IfHostVerified(clientHost string) (bool, error) {
	ctx := context.Background()
	filter := bson.M{
		"host": clientHost,
	}
	hostsResponse, err := dc.engine.Database(dc.dbName).Collection("host").Find(ctx, filter)
	if err != nil {
		return false, err
	}

	defer hostsResponse.Close(ctx)

	var hosts []model.Host
	for hostsResponse.Next(ctx) {
		var host model.Host
		err = hostsResponse.Decode(&host)
		if err != nil {
			return false, err
		}
		hosts = append(hosts, host)
	}

	if len(hosts) > 0 {
		return true, nil
	}
	return false, nil
}

func (dc *DaoClient) IfHostInThisProjectAlreadyExist(host model.Host) (bool, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": host.ProjectId,
		"host":       host.Host,
		"deleted_at": nil,
	}
	var searchedHost model.Host
	err := dc.engine.Database(dc.dbName).Collection("host").FindOne(ctx, filter).Decode(&searchedHost)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (dc *DaoClient) CreateHost(host model.Host) (*model.Host, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": host.ProjectId,
		"host":       host.Host,
		"deleted_at": nil,
	}
	var searchedHost *model.Host
	err := dc.engine.Database(dc.dbName).Collection("host").FindOne(ctx, filter).Decode(&searchedHost)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if searchedHost != nil {
		return nil, fmt.Errorf("host with this name already exist")
	}

	now := time.Now().UTC()
	host.CreatedAt = &now
	host.UpdatedAt = &now
	_, err = dc.engine.Database(dc.dbName).Collection("host").InsertOne(ctx, host)
	if err != nil {
		return nil, err
	}

	var responseHost *model.Host
	afterInsertFilter := bson.M{
		"project_id": host.ProjectId,
		"host":       host.Host,
		"deleted_at": nil,
	}
	err = dc.engine.Database(dc.dbName).Collection("host").FindOne(ctx, afterInsertFilter).Decode(&responseHost)
	if err != nil {
		return nil, err
	}

	return responseHost, nil
}

func (dc *DaoClient) GetHost(hostFilter model.Host) (*model.Host, error) {
	ctx := context.Background()
	filter := bson.M{
		"_id":        hostFilter.Id,
		"project_id": hostFilter.ProjectId,
		"deleted_at": nil,
	}
	var searchedHost *model.Host
	err := dc.engine.Database(dc.dbName).Collection("host").FindOne(ctx, filter).Decode(&searchedHost)
	return searchedHost, err
}

func (dc *DaoClient) ListHosts(hostFilter model.Host) ([]*model.Host, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": hostFilter.ProjectId,
		"deleted_at": nil,
	}
	searchedHostResponse, err := dc.engine.Database(dc.dbName).Collection("host").Find(ctx, filter)

	defer searchedHostResponse.Close(ctx)

	var hosts []*model.Host
	for searchedHostResponse.Next(ctx) {
		var host model.Host
		if err = searchedHostResponse.Decode(&host); err != nil {
			return nil, err
		}
		hosts = append(hosts, &host)
	}

	return hosts, err
}

func (dc *DaoClient) UpdateHost(host model.Host) (*model.Host, error) {
	ctx := context.Background()
	filter := bson.M{
		"_id":        host.Id,
		"project_id": host.ProjectId,
		"deleted_at": nil,
	}

	update := bson.D{{"$set", bson.D{
		{"verified", host.Verified},
		{"updated_at", time.Now().UTC()},
	}}}

	_, err := dc.engine.Database(dc.dbName).Collection("host").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var responseHost model.Host
	err = dc.engine.Database(dc.dbName).Collection("host").FindOne(ctx, filter).Decode(&responseHost)
	if err != nil {
		return nil, err
	}

	return &responseHost, nil
}

func (dc *DaoClient) DeleteHost(hostFilter model.Host) error {
	ctx := context.Background()
	filter := bson.M{
		"_id":        hostFilter.Id,
		"project_id": hostFilter.ProjectId,
		"deleted_at": nil,
	}

	update := bson.D{{"$set", bson.D{
		{"deleted_at", time.Now().UTC()},
	}}}

	_, err := dc.engine.Database(dc.dbName).Collection("host").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
