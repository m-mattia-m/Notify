package service

import "notify/internal/model"

func (dbc *DbClient) GetCredential(credentials model.Credential) (*model.Credential, error) {
	return dbc.dao.GetCredential(credentials)
}

func (dbc *DbClient) ListCredential(credentialFilter model.Credential) ([]*model.Credential, error) {
	return dbc.dao.ListCredential(credentialFilter)
}
