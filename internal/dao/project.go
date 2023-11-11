package dao

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"message-proxy/internal/model"
	"time"
)

func (dc *DaoClient) IfProjectWithThisNameAlreadyExist(project model.Project) (bool, error) {
	ctx := context.Background()
	filter := bson.M{
		"user_id":    project.UserId,
		"name":       project.Name,
		"deleted_at": nil,
	}
	var searchedProject model.Project
	err := dc.engine.Database(dc.dbName).Collection("project").FindOne(ctx, filter).Decode(&searchedProject)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (dc *DaoClient) CreateProject(project model.Project) (*model.Project, error) {
	ctx := context.Background()
	filter := bson.M{
		"user_id":    project.UserId,
		"name":       project.Name,
		"deleted_at": nil,
	}
	var searchedProject *model.Project
	err := dc.engine.Database(dc.dbName).Collection("project").FindOne(ctx, filter).Decode(&searchedProject)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if searchedProject != nil {
		return nil, fmt.Errorf("project with this name already exist")
	}

	now := time.Now().UTC()
	project.CreatedAt = &now
	project.UpdatedAt = &now
	_, err = dc.engine.Database(dc.dbName).Collection("project").InsertOne(ctx, project)
	if err != nil {
		return nil, err
	}

	var responseProject model.Project
	err = dc.engine.Database(dc.dbName).Collection("project").FindOne(ctx, filter).Decode(&responseProject)
	if err != nil {
		return nil, err
	}

	return &responseProject, nil
}

func (dc *DaoClient) GetProject(projectFilter model.Project) (*model.Project, error) {
	ctx := context.Background()
	filter := bson.M{
		"_id":        projectFilter.Id,
		"user_id":    projectFilter.UserId,
		"deleted_at": nil,
	}
	var searchedProject *model.Project
	err := dc.engine.Database(dc.dbName).Collection("project").FindOne(ctx, filter).Decode(&searchedProject)
	return searchedProject, err
}

func (dc *DaoClient) ListProjects(projectFilter model.Project) ([]*model.Project, error) {
	ctx := context.Background()
	filter := bson.M{
		"user_id":    projectFilter.UserId,
		"deleted_at": nil,
	}
	searchedProjectsResponse, err := dc.engine.Database(dc.dbName).Collection("project").Find(ctx, filter)

	defer searchedProjectsResponse.Close(ctx)

	var projects []*model.Project
	for searchedProjectsResponse.Next(ctx) {
		var project model.Project
		if err = searchedProjectsResponse.Decode(&project); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, err
}

func (dc *DaoClient) UpdateProject(project model.Project) (*model.Project, error) {
	ctx := context.Background()
	filter := bson.M{
		"_id":        project.Id,
		"user_id":    project.UserId,
		"deleted_at": nil,
	}

	update := bson.D{{"$set", bson.D{
		{"name", project.Name},
		{"updated_at", time.Now().UTC()},
	}}}

	_, err := dc.engine.Database(dc.dbName).Collection("project").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var responseProject model.Project
	err = dc.engine.Database(dc.dbName).Collection("project").FindOne(ctx, filter).Decode(&responseProject)
	if err != nil {
		return nil, err
	}

	return &responseProject, nil
}

func (dc *DaoClient) DeleteProject(projectFilter model.Project) error {
	ctx := context.Background()
	filter := bson.M{
		"_id":        projectFilter.Id,
		"user_id":    projectFilter.UserId,
		"deleted_at": nil,
	}

	update := bson.D{{"$set", bson.D{
		{"deleted_at", time.Now().UTC()},
	}}}

	_, err := dc.engine.Database(dc.dbName).Collection("project").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
