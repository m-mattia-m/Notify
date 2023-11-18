package dao

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"notify/internal/model"
	"time"
)

func (dc *DaoClient) IfFlowInThisProjectAlreadyExist(flowFilter model.Flow) (bool, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": flowFilter.ProjectId,
		"name":       flowFilter.Name,
		"deleted_at": nil,
	}
	var searchedFlow model.Flow
	err := dc.engine.Database(dc.dbName).Collection("flow").FindOne(ctx, filter).Decode(&searchedFlow)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (dc *DaoClient) CreateFlow(flow model.Flow) (*model.Flow, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": flow.ProjectId,
		"name":       flow.Name,
		"deleted_at": nil,
	}
	var searchedFlow *model.Flow
	err := dc.engine.Database(dc.dbName).Collection("flow").FindOne(ctx, filter).Decode(&searchedFlow)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if searchedFlow != nil {
		return nil, fmt.Errorf("flow with this name already exist")
	}

	now := time.Now().UTC()
	flow.CreatedAt = &now
	flow.UpdatedAt = &now
	_, err = dc.engine.Database(dc.dbName).Collection("flow").InsertOne(ctx, flow)
	if err != nil {
		return nil, err
	}

	var responseFlow model.Flow
	err = dc.engine.Database(dc.dbName).Collection("flow").FindOne(ctx, filter).Decode(&responseFlow)
	if err != nil {
		return nil, err
	}

	return &responseFlow, nil
}

func (dc *DaoClient) GetFlow(flowFilter model.Flow) (*model.Flow, error) {
	ctx := context.Background()
	filter := bson.M{
		"_id":        flowFilter.Id,
		"project_id": flowFilter.ProjectId,
		"deleted_at": nil,
	}
	var searchedFlow *model.Flow
	err := dc.engine.Database(dc.dbName).Collection("flow").FindOne(ctx, filter).Decode(&searchedFlow)
	return searchedFlow, err
}

func (dc *DaoClient) ListFlows(flowFilter model.Flow) ([]*model.Flow, error) {
	ctx := context.Background()
	filter := bson.M{
		"project_id": flowFilter.ProjectId,
		"deleted_at": nil,
	}
	searchedFlowsResponse, err := dc.engine.Database(dc.dbName).Collection("flow").Find(ctx, filter)

	defer searchedFlowsResponse.Close(ctx)

	var flows []*model.Flow
	for searchedFlowsResponse.Next(ctx) {
		var flow model.Flow
		if err = searchedFlowsResponse.Decode(&flow); err != nil {
			return nil, err
		}
		flows = append(flows, &flow)
	}

	return flows, err
}

func (dc *DaoClient) UpdateFlow(flow model.Flow) (*model.Flow, error) {
	ctx := context.Background()
	filter := bson.M{
		"_id":        flow.Id,
		"project_id": flow.ProjectId,
		"deleted_at": nil,
	}

	var searchedFlow *model.Flow
	err := dc.engine.Database(dc.dbName).Collection("flow").FindOne(ctx, filter).Decode(&searchedFlow)
	if err != nil {
		return nil, err
	}

	if flow.Name != "" {
		searchedFlow.Name = flow.Name
	}
	if flow.SourceType != "" {
		searchedFlow.SourceType = flow.SourceType
	}
	if flow.Target != "" {
		searchedFlow.Target = flow.Target
	}
	if flow.OverrideTarget != searchedFlow.OverrideTarget {
		searchedFlow.OverrideTarget = flow.OverrideTarget
	}
	if flow.MessageTemplate != "" {
		searchedFlow.MessageTemplate = flow.MessageTemplate
	}
	if flow.MessageTemplateType != "" {
		searchedFlow.MessageTemplateType = flow.MessageTemplateType
	}
	if flow.Active != flow.Active {
		searchedFlow.Active = flow.Active
	}

	update := bson.D{{"$set", bson.D{
		{"name", searchedFlow.Name},
		{"source_type", searchedFlow.SourceType},
		{"target", searchedFlow.Target},
		{"override_target", searchedFlow.OverrideTarget},
		{"message_template", searchedFlow.MessageTemplate},
		{"message_template_type", searchedFlow.MessageTemplateType},
		{"active", searchedFlow.Active},
		{"updated_at", time.Now().UTC()},
	}}}

	_, err = dc.engine.Database(dc.dbName).Collection("flow").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var responseFlow model.Flow
	err = dc.engine.Database(dc.dbName).Collection("flow").FindOne(ctx, filter).Decode(&responseFlow)
	if err != nil {
		return nil, err
	}

	return &responseFlow, nil
}

func (dc *DaoClient) DeleteFlow(flowFilter model.Flow) error {
	ctx := context.Background()
	filter := bson.M{
		"_id":        flowFilter.Id,
		"project_id": flowFilter.ProjectId,
		"deleted_at": nil,
	}

	update := bson.D{{"$set", bson.D{
		{"deleted_at", time.Now().UTC()},
	}}}

	_, err := dc.engine.Database(dc.dbName).Collection("flow").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
