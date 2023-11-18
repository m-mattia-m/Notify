package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Host struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectId   string             `json:"project_id" bson:"project_id"`
	Host        string             `json:"host" bson:"host"`
	Stage       string             `json:"stage" bson:"stage"` // DEV, TEST, ABT, PROD -> user can create own stages
	Verified    bool               `json:"verified" bson:"verified"`
	VerifyToken string             `json:"verify_token" bson:"verify_token"`
	UpdatedAt   *time.Time         `json:"updated_at" bson:"updated_at"`
	CreatedAt   *time.Time         `json:"created_at" bson:"created_at"`
	DeletedAt   *time.Time         `json:"deleted_at" bson:"deleted_at"`
}

type HostRequest struct {
	Host  string `json:"host"`
	Stage string `json:"stage"`
}
