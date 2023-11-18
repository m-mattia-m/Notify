package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Activity struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectId  string             `json:"project_id" bson:"project_id"`
	State      string             `json:"state" bson:"state"`             // success, failed
	SourceType string             `json:"source_type" bson:"source_type"` // slack, mailgun
	Target     string             `json:"target" bson:"target"`           // e.g. slack-channel-id, to-email
	Subject    string             `json:"subject" bson:"subject"`         // notification subject -> can be enabled/disabled in config.yaml
	Message    string             `json:"message" bson:"message"`         // notification message -> can be enabled/disabled in config.yaml
	Note       string             `json:"note" bson:"note"`               // administration/technical information
	UpdatedAt  *time.Time         `json:"updated_at" bson:"updated_at"`
	CreatedAt  *time.Time         `json:"created_at" bson:"created_at"`
	DeletedAt  *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
