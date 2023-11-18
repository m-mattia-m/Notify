package service

import "notify/internal/model"

func (dbc *DbClient) CreateSlackCredential(credentials model.SlackCredentials) error {
	return dbc.dao.CreateSlackCredential(credentials)
}

func (dbc *DbClient) IsSlackCredentialsAlreadySet(credentials model.SlackCredentials) (bool, error) {
	return dbc.dao.IsSlackCredentialsAlreadySet(credentials)
}

func (dbc *DbClient) UpdateSlackCredential(credentials model.SlackCredentials) error {
	return dbc.dao.UpdateSlackCredential(credentials)
}

func (dbc *DbClient) DeleteSlackCredential(credentials model.SlackCredentials) error {
	return dbc.dao.DeleteSlackCredential(credentials)
}

func (dbc *DbClient) GetSlackRevealedCredential(credentialsFilter model.SlackCredentials) (*model.SlackCredentials, error) {
	return dbc.dao.GetSlackRevealedCredential(credentialsFilter)
}
