package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"notify/internal/helper"
	"notify/internal/model"
)

// CreateFlow			 	godoc
// @title           		CreateFlow
// @description     		Create a flow for notifications. In the flow you can define which host should trigger which notifications and who is the default recipient. You can also define a message template as to how the message should look by default (you can replace data dynamically).
// @Tags 					Flow
// @Router  				/settings/projects/{projectId}/flows [post]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    		path     	string  				true  	"projectId"
// @Param					FlowRequest 		body 		model.FlowRequest 		true 	"FlowRequest"
// @Success      			200  				{object} 	model.Flow
// @Failure      			400  				{object} 	model.HttpError
// @Failure      			404  				{object} 	model.HttpError
// @Failure      			500  				{object} 	model.HttpError
func (sac *SettingsApiClient) CreateFlow(c *gin.Context) {
	svc, _, err := getServiceAndUser(c, false)
	if err != nil {
		log.Error("failed to get service or user from context :: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "an unexpected error has occurred"})
		return
	}

	projectId := c.Param("projectId")
	if projectId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "projectId is required"})
		return
	}

	var flowRequest model.FlowRequest
	err = c.BindJSON(&flowRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to flow-request-object"})
		return
	}

	flow, err := svc.CreateFlow(projectId, flowRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to create flow."),
		})
		return
	}

	c.JSON(http.StatusOK, flow)
}

// ListFlow			 		godoc
// @title           		ListFlow
// @description     		List all Flows in a project
// @Tags 					Flow
// @Router  				/settings/projects/{projectId}/flows [get]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    		path     	string  				true  	"projectId"
// @Success      			200  				{object} 	[]model.Flow
// @Failure      			400  				{object} 	model.HttpError
// @Failure      			404  				{object} 	model.HttpError
// @Failure      			500  				{object} 	model.HttpError
func (sac *SettingsApiClient) ListFlow(c *gin.Context) {
	svc, _, err := getServiceAndUser(c, false)
	if err != nil {
		log.Error("failed to get service or user from context :: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "an unexpected error has occurred"})
		return
	}

	projectId := c.Param("projectId")
	if projectId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "projectId is required"})
		return
	}

	flows, err := svc.ListFlows(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to list flows."),
		})
		return
	}

	c.JSON(http.StatusOK, flows)
}

// GetFlow			 		godoc
// @title           		GetFlow
// @description     		Get a specific flow by their id
// @Tags 					Flow
// @Router  				/settings/projects/{projectId}/flows/{flowId} [get]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    		path     	string  				true  	"projectId"
// @Param        			flowId	    		path     	string  				true  	"flowId"
// @Success      			200  				{object} 	model.Flow
// @Failure      			400  				{object} 	model.HttpError
// @Failure      			404  				{object} 	model.HttpError
// @Failure      			500  				{object} 	model.HttpError
func (sac *SettingsApiClient) GetFlow(c *gin.Context) {
	svc, _, err := getServiceAndUser(c, false)
	if err != nil {
		log.Error("failed to get service or user from context :: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "an unexpected error has occurred"})
		return
	}

	projectId := c.Param("projectId")
	if projectId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "projectId is required"})
		return
	}

	flowId := c.Param("flowId")
	if flowId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "flowId is required"})
		return
	}

	flow, err := svc.GetFlow(flowId, projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to create flow."),
		})
		return
	}

	c.JSON(http.StatusOK, flow)
}

// UpdateFlow			 	godoc
// @title           		UpdateFlow
// @description     		Update a specific flow by their id.
// @Tags 					Flow
// @Router  				/settings/projects/{projectId}/flows/{flowId} [put]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    		path     	string  				true  	"projectId"
// @Param        			flowId	    		path     	string  				true  	"flowId"
// @Param					FlowRequest 		body 		model.FlowRequest 		true 	"FlowRequest"
// @Success      			200  				{object} 	model.Flow
// @Failure      			400  				{object} 	model.HttpError
// @Failure      			404  				{object} 	model.HttpError
// @Failure      			500  				{object} 	model.HttpError
func (sac *SettingsApiClient) UpdateFlow(c *gin.Context) {
	svc, _, err := getServiceAndUser(c, false)
	if err != nil {
		log.Error("failed to get service or user from context :: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "an unexpected error has occurred"})
		return
	}

	projectId := c.Param("projectId")
	if projectId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "projectId is required"})
		return
	}

	flowId := c.Param("flowId")
	if projectId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "flowId is required"})
		return
	}

	var flowRequest model.FlowRequest
	err = c.BindJSON(&flowRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to flow-request-object"})
		return
	}

	flow, err := svc.UpdateFlow(flowId, projectId, flowRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to create flow."),
		})
		return
	}

	c.JSON(http.StatusOK, flow)
}

// DeleteFlow			 	godoc
// @title           		DeleteFlow
// @description     		Delete a specific flow by their id.
// @Tags 					Flow
// @Router  				/settings/projects/{projectId}/flows/{flowId} [delete]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    		path     	string  				true  	"projectId"
// @Param        			flowId    			path     	string  				true  	"flowId"
// @Success      			200  				{object} 	model.Flow
// @Failure      			400  				{object} 	model.HttpError
// @Failure      			404  				{object} 	model.HttpError
// @Failure      			500  				{object} 	model.HttpError
func (sac *SettingsApiClient) DeleteFlow(c *gin.Context) {
	svc, _, err := getServiceAndUser(c, false)
	if err != nil {
		log.Error("failed to get service or user from context :: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "an unexpected error has occurred"})
		return
	}

	projectId := c.Param("projectId")
	if projectId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "projectId is required"})
		return
	}

	flowId := c.Param("flowId")
	if projectId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "flowId is required"})
		return
	}

	flow, err := svc.DeleteFlow(flowId, projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to create flow."),
		})
		return
	}

	c.JSON(http.StatusOK, flow)
}
