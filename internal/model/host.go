package model

type Host struct {
	Id          string `json:"id"`
	ProjectId   string `json:"project_id"`
	Host        string `json:"host"`
	Stage       string `json:"stage"` // DEV, TEST, ABT, PROD -> user can create own stages
	Verified    bool   `json:"verified"`
	VerifyToken string `json:"verify_token"`
}
