package service

import "message-proxy/internal/model"

func (dbc *DbClient) IfProjectWithThisNameAlreadyExist(project model.Project) (bool, error) {
	return dbc.dao.IfProjectWithThisNameAlreadyExist(project)
}

func (dbc *DbClient) CreateProject(project model.Project) (*model.Project, error) {
	return dbc.dao.CreateProject(project)
}

func (dbc *DbClient) GetProject(projectFilter model.Project) (*model.Project, error) {
	return dbc.dao.GetProject(projectFilter)
}

func (dbc *DbClient) ListProjects(projectFilter model.Project) ([]*model.Project, error) {
	return dbc.dao.ListProjects(projectFilter)
}

func (dbc *DbClient) UpdateProject(projectFilter model.Project) (*model.Project, error) {
	return dbc.dao.UpdateProject(projectFilter)
}

func (dbc *DbClient) DeleteProject(projectFilter model.Project) error {
	return dbc.dao.DeleteProject(projectFilter)
}
