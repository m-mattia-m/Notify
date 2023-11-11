package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Project struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	UserId    string             `json:"user_id" bson:"user_id"`
	UpdatedAt *time.Time         `json:"updated_at" bson:"updated_at"`
	CreatedAt *time.Time         `json:"created_at" bson:"created_at"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
}

type ProjectRequest struct {
	Name string `json:"name"`
}
