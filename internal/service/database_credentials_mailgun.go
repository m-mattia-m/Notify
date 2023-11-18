package service

import "notify/internal/model"

func (dbc *DbClient) CreateMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error) {
	return dbc.dao.CreateMailgunCredential(credentials)
}

func (dbc *DbClient) GetMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error) {
	return dbc.dao.GetMailgunCredential(credentials)
}

func (dbc *DbClient) IsMailgunCredentialsAlreadySet(credentials model.MailgunCredentials) (bool, error) {
	return dbc.dao.IsMailgunCredentialsAlreadySet(credentials)
}

func (dbc *DbClient) UpdateMailgunCredential(credentials model.MailgunCredentials) (*model.MailgunCredentialsResponse, error) {
	return dbc.dao.UpdateMailgunCredential(credentials)
}

func (dbc *DbClient) DeleteMailgunCredential(credentials model.MailgunCredentials) error {
	return dbc.dao.DeleteMailgunCredential(credentials)
}

func (dbc *DbClient) GetMailgunRevealedCredential(credentials model.MailgunCredentials) (*model.MailgunCredentials, error) {
	return dbc.dao.GetMailgunRevealedCredential(credentials)
}
