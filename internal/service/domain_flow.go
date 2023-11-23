package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"notify/internal/model"
)

func (c *Client) CreateFlow(projectId string, flowRequest model.FlowRequest) (*model.Flow, error) {
	err := validateFlowRequest(flowRequest)
	if err != nil {
		return nil, err
	}

	flow := model.Flow{
		Name:                flowRequest.Name,
		ProjectId:           projectId,
		SourceType:          flowRequest.SourceType,
		Target:              flowRequest.Target,
		OverrideTarget:      flowRequest.OverrideTarget,
		MessageTemplate:     flowRequest.MessageTemplate,
		MessageTemplateType: flowRequest.MessageTemplateType,
		Active:              flowRequest.Active,
	}

	if !c.proveMessageTypes(flowRequest.MessageTemplateType) {
		return nil, fmt.Errorf("invalid message template type")
	}

	alreadyExist, err := c.db.IfFlowInThisProjectAlreadyExist(flow)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	if alreadyExist {
		return nil, fmt.Errorf("flow with this name already exist in this project")
	}

	flowResponse, err := c.db.CreateFlow(flow)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return flowResponse, nil
}

func (c *Client) GetFlow(flowId string, projectId string) (*model.Flow, error) {
	flowObjectId, err := primitive.ObjectIDFromHex(flowId)
	if err != nil {
		return nil, fmt.Errorf("invalid flowId")
	}

	projectFilter := model.Flow{
		Id:        flowObjectId,
		ProjectId: projectId,
	}
	projectResponse, err := c.db.GetFlow(projectFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return projectResponse, nil
}

func (c *Client) ListFlows(projectId string) ([]*model.Flow, error) {
	flowFilter := model.Flow{
		ProjectId: projectId,
	}

	flowListResponse, err := c.db.ListFlows(flowFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return flowListResponse, nil
}

func (c *Client) UpdateFlow(flowId, projectId string, flowRequest model.FlowRequest) (*model.Flow, error) {
	err := validateFlowRequest(flowRequest)
	if err != nil {
		return nil, err
	}

	flowObjectId, err := primitive.ObjectIDFromHex(flowId)
	if err != nil {
		return nil, fmt.Errorf("invalid flowId")
	}

	if !c.proveMessageTypes(flowRequest.MessageTemplateType) {
		return nil, fmt.Errorf("invalid message template type")
	}

	flowFilter := model.Flow{
		Id:                  flowObjectId,
		ProjectId:           projectId,
		Name:                flowRequest.Name,
		SourceType:          flowRequest.SourceType,
		Target:              flowRequest.Target,
		OverrideTarget:      flowRequest.OverrideTarget,
		MessageTemplate:     flowRequest.MessageTemplate,
		MessageTemplateType: flowRequest.MessageTemplateType,
		Active:              flowRequest.Active,
	}

	flowResponse, err := c.db.UpdateFlow(flowFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return flowResponse, nil
}

func (c *Client) DeleteFlow(flowId string, projectId string) (*model.SuccessMessage, error) {
	flowObjectId, err := primitive.ObjectIDFromHex(flowId)
	if err != nil {
		return nil, fmt.Errorf("invalid flowId")
	}

	flowFilter := model.Flow{
		Id:        flowObjectId,
		ProjectId: projectId,
	}
	flowResponse, err := c.db.GetFlow(flowFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	err = c.db.DeleteFlow(flowFilter)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("")
	}

	return &model.SuccessMessage{
		Message: fmt.Sprintf("%s successfully deleted", flowResponse.Name),
	}, nil
}

func (c *Client) proveMessageTypes(messageType string) bool {
	switch messageType {
	case "text/plain":
		return true
	case "text/html":
		return true
	default:
		return false
	}
}

func validateFlowRequest(flowRequest model.FlowRequest) error {
	if flowRequest.Name == "" {
		return fmt.Errorf("name is a required attribute")
	}
	if flowRequest.SourceType == "" {
		return fmt.Errorf("name is a required attribute")
	}
	if flowRequest.Target == "" {
		return fmt.Errorf("name is a required attribute")
	}
	if flowRequest.MessageTemplateType == "text/plain" || flowRequest.MessageTemplateType == "text/html" {
		return fmt.Errorf("message-template-type is a required attribute; only 'text/plain' or 'text/html' are valid")
	}

	return nil
}
