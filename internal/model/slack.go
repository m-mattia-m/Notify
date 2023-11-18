package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SlackCredentials struct {
	Id                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type              string             `json:"-" bson:"type"`
	ProjectId         string             `json:"project_id" bson:"project_id"`
	BotUserOAuthToken string             `json:"bot_user_oauth_token" bson:"bot_user_oauth_token"`
	UpdatedAt         *time.Time         `json:"updated_at" bson:"updated_at"`
	CreatedAt         *time.Time         `json:"created_at" bson:"created_at"`
}

type SlackCredentialsRequest struct {
	BotUserOAuthToken string `json:"bot_user_o_auth_token"`
}
