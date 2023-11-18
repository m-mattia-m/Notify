package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"notify/internal/model"
)

func (c *Client) GetActivity(activityId, projectId string) (*model.Activity, error) {
	activityObjectId, err := primitive.ObjectIDFromHex(activityId)
	if err != nil {
		return nil, fmt.Errorf("invalid activityId")
	}

	return c.db.GetActivity(model.Activity{
		Id:        activityObjectId,
		ProjectId: projectId,
	})
}

func (c *Client) ListActivities(projectId string) ([]*model.Activity, error) {
	return c.db.ListActivities(model.Activity{ProjectId: projectId})
}
