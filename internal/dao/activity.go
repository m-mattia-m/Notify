package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"notify/internal/model"
	"time"
)

func (dc *DaoClient) CreateActivity(activity model.Activity) error {
	ctx := context.Background()
	now := time.Now().UTC()
	activity.CreatedAt = &now
	activity.UpdatedAt = &now
	_, err := dc.engine.Database(dc.dbName).Collection("activity").InsertOne(ctx, activity)
	return err
}

func (dc *DaoClient) GetActivity(activityFilter model.Activity) (*model.Activity, error) {
	ctx := context.Background()
	filter := bson.M{
		"_id":        activityFilter.Id,
		"project_id": activityFilter.ProjectId,
		"deleted_at": nil,
	}
	var searchedActivity *model.Activity
	err := dc.engine.Database(dc.dbName).Collection("activity").FindOne(ctx, filter).Decode(&searchedActivity)
	return searchedActivity, err
}

func (dc *DaoClient) ListActivities(activityFilter model.Activity) ([]*model.Activity, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": activityFilter.ProjectId,
		"deleted_at": nil,
	}
	searchedActivityResponse, err := dc.engine.Database(dc.dbName).Collection("activity").Find(ctx, filter)

	defer searchedActivityResponse.Close(ctx)

	var activities []*model.Activity
	for searchedActivityResponse.Next(ctx) {
		var activity model.Activity
		if err = searchedActivityResponse.Decode(&activity); err != nil {
			return nil, err
		}
		activities = append(activities, &activity)
	}

	return activities, err
}
