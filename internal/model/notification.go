package model

type Notification struct {
	ProjectId string `json:"project_id"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	Target    string `json:"target"`
}
