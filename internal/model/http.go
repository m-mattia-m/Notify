package model

type HttpError struct {
	Message string `json:"message"`
}

type SuccessMessage struct {
	Message string `json:"message"`
}

type StateResponse struct {
	State bool `json:"state"`
}
