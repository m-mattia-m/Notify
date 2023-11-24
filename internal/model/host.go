package model

import (
	"fmt"
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

func (h *Host) String() string {
	return fmt.Sprintf(
		fmt.Sprintf("Id: %s", h.Id.String()),
		fmt.Sprintf("ProjectId: %s", h.ProjectId),
		fmt.Sprintf("Host: %s", h.Host),
		fmt.Sprintf("Stage: %s", h.Stage),
		fmt.Sprintf("Verified: %s", h.VerifyToken),
		fmt.Sprintf("VerifyToken: %s", h.VerifyToken),
		fmt.Sprintf("UpdatedAt: %s", h.UpdatedAt),
		fmt.Sprintf("CreatedAt: %s", h.CreatedAt),
		fmt.Sprintf("DeletedAt: %s", h.DeletedAt),
	)
}
