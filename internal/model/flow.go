package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Flow struct {
	Id                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name                string             `json:"name" bson:"name"`                                   // e.g. company notification (slack) || sender notification (mailgun)
	ProjectId           string             `json:"project_id" bson:"project_id"`                       // the projectId a flow belongs to
	SourceType          string             `json:"source_type" bson:"source_type"`                     // message provider via which the message is sent || keyword: mailgun
	Target              string             `json:"target" bson:"target"`                               // can be overwritten in the request. e.g. with target-Email -> slack-channel-id || override: sender-email
	OverrideTarget      bool               `json:"override_target" bson:"override_target"`             // if true, then override e.g. the configured email with the email from the request
	MessageTemplate     string             `json:"message_template" bson:"message_template"`           // layout with the message -> Title: \{\{ notification.Subject \}\} \nCustomer creates new support-ticket via frontend-form with the message: \{\{ notification.message \}\}.
	MessageTemplateType string             `json:"message_template_type" bson:"message_template_type"` // defines if the template is a TXT or HTML, Markdown, ...
	Active              bool               `json:"active" bson:"active"`                               // defines if the workflow should be triggered
	UpdatedAt           *time.Time         `json:"updated_at" bson:"updated_at"`
	CreatedAt           *time.Time         `json:"created_at" bson:"created_at"`
	DeletedAt           *time.Time         `json:"deleted_at" bson:"deleted_at"`
}

type FlowRequest struct {
	Name                string `json:"name" bson:"name"`                                   // e.g. company notification (slack) || sender notification (mailgun)
	SourceType          string `json:"source_type" bson:"source_type"`                     // message provider via which the message is sent || keyword: mailgun
	Target              string `json:"target" bson:"target"`                               // can be overwritten in the request. e.g. with target-Email -> slack-channel-id || override: sender-email
	OverrideTarget      bool   `json:"override_target" bson:"override_target"`             // if true, then override e.g. the configured email with the email from the request
	MessageTemplate     string `json:"message_template" bson:"message_template"`           // layout with the message -> Title: \{\{ notification.Subject \}\} \nCustomer creates new support-ticket via frontend-form with the message: \{\{ notification.message \}\}.
	MessageTemplateType string `json:"message_template_type" bson:"message_template_type"` // defines if the template is a TXT or HTML, Markdown, ...
	Active              bool   `json:"active" bson:"active"`                               // defines if the workflow should be triggered
}
