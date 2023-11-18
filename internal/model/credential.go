package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credential struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectId string             `json:"project_id" bson:"project_id"`
	Type      string             `json:"-" bson:"type"`
}
