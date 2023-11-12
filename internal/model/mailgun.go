package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MailgunCredentials struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type        string             `json:"-" bson:"type"`
	ProjectId   string             `json:"project_id" bson:"project_id"`
	Domain      string             `json:"domain" bson:"domain"`
	ApiKey      string             `json:"api_key" bson:"api_key"`
	SenderEmail string             `json:"sender_email" bson:"sender_email"`
	SenderName  string             `json:"sender_name" bson:"sender_name"`
	UpdatedAt   *time.Time         `json:"updated_at" bson:"updated_at"`
	CreatedAt   *time.Time         `json:"created_at" bson:"created_at"`
}

type MailgunCredentialsRequest struct {
	Domain      string `json:"domain"`
	ApiKey      string `json:"api_key"`
	SenderEmail string `json:"sender_email"`
	SenderName  string `json:"sender_name"`
}

type MailgunCredentialsResponse struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectId   string             `json:"project_id" bson:"project_id"`
	Domain      string             `json:"domain"`
	SenderEmail string             `json:"sender_email"`
	SenderName  string             `json:"sender_name"`
	CreatedAt   *time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time         `json:"updated_at" bson:"updated_at"`
}
