package service

import "notify/internal/model"

func (dbc *DbClient) CreateActivity(activity model.Activity) error {
	return dbc.dao.CreateActivity(activity)
}

func (dbc *DbClient) GetActivity(activityFilter model.Activity) (*model.Activity, error) {
	return dbc.dao.GetActivity(activityFilter)
}

func (dbc *DbClient) ListActivities(activityFilter model.Activity) ([]*model.Activity, error) {
	return dbc.dao.ListActivities(activityFilter)
}
