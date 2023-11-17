package service

import "notify/internal/model"

func (dbc *DbClient) IfFlowInThisProjectAlreadyExist(flowFilter model.Flow) (bool, error) {
	return dbc.dao.IfFlowInThisProjectAlreadyExist(flowFilter)
}

func (dbc *DbClient) CreateFlow(flow model.Flow) (*model.Flow, error) {
	return dbc.dao.CreateFlow(flow)
}

func (dbc *DbClient) GetFlow(flowFilter model.Flow) (*model.Flow, error) {
	return dbc.dao.GetFlow(flowFilter)
}

func (dbc *DbClient) ListFlows(flowFilter model.Flow) ([]*model.Flow, error) {
	return dbc.dao.ListFlows(flowFilter)
}

func (dbc *DbClient) UpdateFlow(flow model.Flow) (*model.Flow, error) {
	return dbc.dao.UpdateFlow(flow)
}

func (dbc *DbClient) DeleteFlow(flowFilter model.Flow) error {
	return dbc.dao.DeleteFlow(flowFilter)
}
