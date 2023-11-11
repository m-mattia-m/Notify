package service

import "message-proxy/internal/model"

func (dbc *DbClient) IfHostVerified(clientHost string) (bool, error) {
	return dbc.dao.IfHostVerified(clientHost)
}

func (dbc *DbClient) IfHostInThisProjectAlreadyExist(host model.Host) (bool, error) {
	return dbc.dao.IfHostInThisProjectAlreadyExist(host)
}

func (dbc *DbClient) CreateHost(host model.Host) (*model.Host, error) {
	return dbc.dao.CreateHost(host)
}

func (dbc *DbClient) GetHost(hostFilter model.Host) (*model.Host, error) {
	return dbc.dao.GetHost(hostFilter)
}

func (dbc *DbClient) ListHosts(hostFilter model.Host) ([]*model.Host, error) {
	return dbc.dao.ListHosts(hostFilter)
}

func (dbc *DbClient) UpdateHost(host model.Host) (*model.Host, error) {
	return dbc.dao.UpdateHost(host)
}

func (dbc *DbClient) DeleteHost(hostFilter model.Host) error {
	return dbc.dao.DeleteHost(hostFilter)
}
