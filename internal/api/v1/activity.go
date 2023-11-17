package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"notify/internal/helper"
)

// ListActivities			godoc
// @title           		ListActivities
// @description     		List all activities in a project
// @Tags 					Activity
// @Router  				/settings/projects/{projectId}/activities [get]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    		path     	string  				true  	"projectId"
// @Success      			200  				{object} 	[]model.Activity
// @Failure      			400  				{object} 	model.HttpError
// @Failure      			404  				{object} 	model.HttpError
// @Failure      			500  				{object} 	model.HttpError
func (sac *SettingsApiClient) ListActivities(c *gin.Context) {
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

	activities, err := svc.ListActivities(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to list activities."),
		})
		return
	}

	c.JSON(http.StatusOK, activities)
}

// GetActivity		 		godoc
// @title           		GetActivity
// @description     		Get a specific activity by their id
// @Tags 					Activity
// @Router  				/settings/projects/{projectId}/activities/{activityId} [get]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    		path     	string  				true  	"projectId"
// @Param        			activityId	    	path     	string  				true  	"activityId"
// @Success      			200  				{object} 	model.Activity
// @Failure      			400  				{object} 	model.HttpError
// @Failure      			404  				{object} 	model.HttpError
// @Failure      			500  				{object} 	model.HttpError
func (sac *SettingsApiClient) GetActivity(c *gin.Context) {
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

	activityId := c.Param("activityId")
	if activityId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "activityId is required"})
		return
	}

	activity, err := svc.GetActivity(activityId, projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to create flow."),
		})
		return
	}

	c.JSON(http.StatusOK, activity)
}
