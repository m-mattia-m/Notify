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

const CREDENTIAL_TYPE_SLACK = "slack"

func (dc *DaoClient) CreateSlackCredential(credentials model.SlackCredentials) error {

	ctx := context.Background()
	filter := bson.M{
		"project_id": credentials.ProjectId,
		"type":       CREDENTIAL_TYPE_SLACK,
	}
	var searchedCredentials *model.SlackCredentials
	err := dc.engine.Database(dc.dbName).Collection("credential").FindOne(ctx, filter).Decode(&searchedCredentials)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	if searchedCredentials != nil {
		return fmt.Errorf("slack credential already exist")
	}

	now := time.Now().UTC()
	credentials.CreatedAt = &now
	credentials.UpdatedAt = &now
	credentials.Type = CREDENTIAL_TYPE_SLACK
	_, err = dc.engine.Database(dc.dbName).Collection("credential").InsertOne(ctx, credentials)
	if err != nil {
		return err
	}

	return nil
}

func (dc *DaoClient) IsSlackCredentialsAlreadySet(credentials model.SlackCredentials) (bool, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": credentials.ProjectId,
		"type":       CREDENTIAL_TYPE_SLACK,
	}
	var searchedCredentials *model.SlackCredentials
	err := dc.engine.Database(dc.dbName).Collection("credential").FindOne(ctx, filter).Decode(&searchedCredentials)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return false, err
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	return true, err
}

func (dc *DaoClient) UpdateSlackCredential(credentials model.SlackCredentials) error {
	ctx := context.Background()
	filter := bson.M{
		// There is currently no way to read out the `_id`, but since the `type` and `project_id` are unique, an object is currently identified via them.
		//"_id":        credentials.Id,
		"project_id": credentials.ProjectId,
		"type":       CREDENTIAL_TYPE_SLACK,
	}

	update := bson.D{{"$set", bson.D{
		{"bot_user_oauth_token", credentials.BotUserOAuthToken},
		{"updated_at", time.Now().UTC()},
	}}}

	_, err := dc.engine.Database(dc.dbName).Collection("credential").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (dc *DaoClient) DeleteSlackCredential(credentials model.SlackCredentials) error {
	ctx := context.Background()
	filter := bson.M{
		"project_id": credentials.ProjectId,
		"type":       CREDENTIAL_TYPE_SLACK,
	}

	// The objects are actually deleted here so that no unnecessary sensitive data continues to exist.
	_, err := dc.engine.Database(dc.dbName).Collection("credential").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (dc *DaoClient) GetSlackRevealedCredential(credentialsFilter model.SlackCredentials) (*model.SlackCredentials, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": credentialsFilter.ProjectId,
		"type":       CREDENTIAL_TYPE_SLACK,
	}
	var searchedCredentials *model.SlackCredentials
	err := dc.engine.Database(dc.dbName).Collection("credential").FindOne(ctx, filter).Decode(&searchedCredentials)
	if err != nil {
		return nil, err
	}

	return searchedCredentials, nil
}
