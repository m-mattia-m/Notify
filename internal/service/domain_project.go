package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"notify/internal/model"
)

func (c *Client) CreateProject(projectRequest model.ProjectRequest, userId string) (*model.Project, error) {
	project := model.Project{
		Name:   projectRequest.Name,
		UserId: userId,
	}

	alreadyExist, err := c.db.IfProjectWithThisNameAlreadyExist(project)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if alreadyExist {
		return nil, fmt.Errorf("project with this name already exist")
	}

	projectResponse, err := c.db.CreateProject(project)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return projectResponse, nil
}

func (c *Client) GetProject(projectId, userId string) (*model.Project, error) {
	projectObjectId, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, fmt.Errorf("invalid projectId")
	}

	projectFilter := model.Project{
		Id:     projectObjectId,
		UserId: userId,
	}

	projectResponse, err := c.db.GetProject(projectFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return projectResponse, nil
}

func (c *Client) ListProjects(userId string) ([]*model.Project, error) {
	projectFilter := model.Project{
		UserId: userId,
	}

	projectListResponse, err := c.db.ListProjects(projectFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return projectListResponse, nil
}

func (c *Client) UpdateProject(projectId string, projectRequest model.ProjectRequest, userId string) (*model.Project, error) {
	projectObjectId, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, fmt.Errorf("invalid projectId")
	}

	projectFilter := model.Project{
		Id:     projectObjectId,
		UserId: userId,
		Name:   projectRequest.Name,
	}

	projectResponse, err := c.db.UpdateProject(projectFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return projectResponse, nil
}

func (c *Client) DeleteProject(projectId, userId string) (*model.SuccessMessage, error) {
	projectObjectId, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, fmt.Errorf("invalid projectId")
	}

	projectFilter := model.Project{
		Id:     projectObjectId,
		UserId: userId,
	}
	projectResponse, err := c.db.GetProject(projectFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	err = c.db.DeleteProject(projectFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return &model.SuccessMessage{
		Message: fmt.Sprintf("%s successfully deleted", projectResponse.Name),
	}, nil
}
