package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"notify/internal/model"
)

func (dc *DaoClient) GetCredential(credentials model.Credential) (*model.Credential, error) {
	ctx := context.Background()
	filter := bson.M{
		"_id":        credentials.Id,
		"project_id": credentials.ProjectId,
	}

	var searchedCredential *model.Credential
	err := dc.engine.Database(dc.dbName).Collection("credential").FindOne(ctx, filter).Decode(&searchedCredential)
	if err != nil {
		return nil, err
	}

	return searchedCredential, nil
}

func (dc *DaoClient) ListCredential(credentialFilter model.Credential) ([]*model.Credential, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": credentialFilter.ProjectId,
	}

	searchedCredentialsResponse, err := dc.engine.Database(dc.dbName).Collection("credential").Find(ctx, filter)
	defer searchedCredentialsResponse.Close(ctx)

	var credentials []*model.Credential
	for searchedCredentialsResponse.Next(ctx) {
		var credential model.Credential
		if err = searchedCredentialsResponse.Decode(&credential); err != nil {
			return nil, err
		}
		credentials = append(credentials, &credential)
	}

	return credentials, nil

}
