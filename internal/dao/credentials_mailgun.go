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

const CREDENTIAL_TYPE_MAILGUN = "mailgun"

func (dc *DaoClient) CreateMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": credentials.ProjectId,
		"type":       CREDENTIAL_TYPE_MAILGUN,
	}
	var searchedCredentials *model.MailgunCredentials
	err := dc.engine.Database(dc.dbName).Collection("credential").FindOne(ctx, filter).Decode(&searchedCredentials)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if searchedCredentials != nil {
		return nil, fmt.Errorf("mailgun credential already exist")
	}

	now := time.Now().UTC()
	credentials.CreatedAt = &now
	credentials.UpdatedAt = &now
	credentials.Type = CREDENTIAL_TYPE_MAILGUN
	_, err = dc.engine.Database(dc.dbName).Collection("credential").InsertOne(ctx, credentials)
	if err != nil {
		return nil, err
	}

	var responseCredentials *model.MailgunCredentials
	err = dc.engine.Database(dc.dbName).Collection("credential").FindOne(ctx, filter).Decode(&responseCredentials)
	if err != nil {
		return nil, err
	}

	return &model.MailgunCredentialsResponse{
		Id:          responseCredentials.Id,
		ProjectId:   responseCredentials.ProjectId,
		Domain:      responseCredentials.Domain,
		SenderEmail: responseCredentials.SenderEmail,
		SenderName:  responseCredentials.SenderName,
		CreatedAt:   responseCredentials.CreatedAt,
		UpdatedAt:   responseCredentials.UpdatedAt,
	}, nil
}

func (dc *DaoClient) GetMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error) {
	ctx := context.Background()
	filter := bson.M{
		// There is currently no way to read out the `_id`, but since the `type` and `project_id` are unique, an object is currently identified via them.
		//"_id":        credentials.Id,
		"project_id": credentials.ProjectId,
		"type":       CREDENTIAL_TYPE_MAILGUN,
	}

	var searchedCredentials *model.MailgunCredentials
	err := dc.engine.Database(dc.dbName).Collection("credential").FindOne(ctx, filter).Decode(&searchedCredentials)
	if err != nil {
		return nil, err
	}

	return &model.MailgunCredentialsResponse{
		Id:          searchedCredentials.Id,
		ProjectId:   searchedCredentials.ProjectId,
		Domain:      searchedCredentials.Domain,
		SenderEmail: searchedCredentials.SenderEmail,
		SenderName:  searchedCredentials.SenderName,
		CreatedAt:   searchedCredentials.CreatedAt,
		UpdatedAt:   searchedCredentials.UpdatedAt,
	}, nil
}

func (dc *DaoClient) IsMailgunCredentialsAlreadySet(credentials model.MailgunCredentials) (bool, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": credentials.ProjectId,
		"type":       CREDENTIAL_TYPE_MAILGUN,
	}
	var searchedCredentials *model.MailgunCredentials
	err := dc.engine.Database(dc.dbName).Collection("credential").FindOne(ctx, filter).Decode(&searchedCredentials)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return false, err
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	return true, err
}

func (dc *DaoClient) UpdateMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error) {
	ctx := context.Background()
	filter := bson.M{
		// There is currently no way to read out the `_id`, but since the `type` and `project_id` are unique, an object is currently identified via them.
		//"_id":        credentials.Id,
		"project_id": credentials.ProjectId,
		"type":       CREDENTIAL_TYPE_MAILGUN,
	}

	var searchedCredentials *model.MailgunCredentials
	err := dc.engine.Database(dc.dbName).Collection("credential").FindOne(ctx, filter).Decode(&searchedCredentials)
	if err != nil {
		return nil, err
	}

	if credentials.Domain != "" {
		searchedCredentials.Domain = credentials.Domain
	}
	if credentials.ApiKey != "" {
		searchedCredentials.ApiKey = credentials.ApiKey
	}
	if credentials.SenderEmail != "" {
		searchedCredentials.SenderEmail = credentials.SenderEmail
	}
	if credentials.SenderName != "" {
		searchedCredentials.SenderName = credentials.SenderName
	}

	updatedTime := time.Now().UTC()
	update := bson.D{{"$set", bson.D{
		{"domain", searchedCredentials.Domain},
		{"api_key", searchedCredentials.ApiKey},
		{"sender_email", searchedCredentials.SenderEmail},
		{"sender_name", searchedCredentials.SenderName},
		{"updated_at", updatedTime},
	}}}

	_, err = dc.engine.Database(dc.dbName).Collection("credential").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &model.MailgunCredentialsResponse{
		Id:          searchedCredentials.Id,
		ProjectId:   searchedCredentials.ProjectId,
		Domain:      searchedCredentials.Domain,
		SenderEmail: searchedCredentials.SenderEmail,
		SenderName:  searchedCredentials.SenderName,
		CreatedAt:   searchedCredentials.CreatedAt,
		UpdatedAt:   searchedCredentials.UpdatedAt,
	}, nil
}

func (dc *DaoClient) DeleteMailgunCredential(credentials model.MailgunCredentials) error {
	ctx := context.Background()
	filter := bson.M{
		"project_id": credentials.ProjectId,
		"type":       CREDENTIAL_TYPE_MAILGUN,
	}

	// The objects are actually deleted here so that no unnecessary sensitive data continues to exist.
	_, err := dc.engine.Database(dc.dbName).Collection("credential").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
