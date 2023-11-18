package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"notify/internal/helper"
	"notify/internal/model"
)

// CreateSlackCredentials 	godoc
// @title           		CreateSlackCredentials
// @description     		Create the access data for your Slack integration.
// @Tags 					IntegrationSlack
// @Router  				/settings/projects/{projectId}/integrations/slack [post]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    				path     	string  						true  	"projectId"
// @Param					SlackCredentialsRequest 	body 		model.SlackCredentialsRequest 	true 	"SlackCredentialsRequest"
// @Success      			200  						{object} 	model.SuccessMessage
// @Failure      			400  						{object} 	model.HttpError
// @Failure      			404  						{object} 	model.HttpError
// @Failure      			500  						{object} 	model.HttpError
func (sac *SettingsApiClient) CreateSlackCredentials(c *gin.Context) {
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

	var slackCredentialRequest model.SlackCredentialsRequest
	err = c.BindJSON(&slackCredentialRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to slack-credentials-request-object"})
		return
	}

	host, err := svc.CreateSlackCredentials(projectId, slackCredentialRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to create slack-credentials."),
		})
		return
	}

	c.JSON(http.StatusOK, host)
}

// UpdateSlackCredentials 	godoc
// @title           		UpdateSlackCredentials
// @description     		Overwrite the token of the Slack access of a project
// @Tags 					IntegrationSlack
// @Router  				/settings/projects/{projectId}/integrations/slack [put]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    				path     	string  						true  	"projectId"
// @Param					SlackCredentialsRequest 	body 		model.SlackCredentialsRequest 	true 	"SlackCredentialsRequest"
// @Success      			200  						{object} 	model.SuccessMessage
// @Failure      			400  						{object} 	model.HttpError
// @Failure      			404  						{object} 	model.HttpError
// @Failure      			500  						{object} 	model.HttpError
func (sac *SettingsApiClient) UpdateSlackCredentials(c *gin.Context) {
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

	var slackCredentialRequest model.SlackCredentialsRequest
	err = c.BindJSON(&slackCredentialRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to slack-credentials-request-object"})
		return
	}

	host, err := svc.UpdateSlackCredentials(projectId, slackCredentialRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to update slack-credentials."),
		})
		return
	}

	c.JSON(http.StatusOK, host)
}

// DeleteSlackCredentials 	godoc
// @title           		DeleteSlackCredentials
// @description     		Deletes slack-credentials from a project
// @Tags 					IntegrationSlack
// @Router  				/settings/projects/{projectId}/integrations/slack [delete]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Success      			200  			{object} 	model.SuccessMessage
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) DeleteSlackCredentials(c *gin.Context) {
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

	successMessage, err := svc.DeleteSlackCredentials(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to delete slack-credentials."),
		})
		return
	}

	c.JSON(http.StatusOK, successMessage)
}

// IsSlackCredentialsAlreadySet		godoc
// @title           				IsSlackCredentialsAlreadySet
// @description     				Indicates whether Slack has already been implemented in a project.
// @Tags 							IntegrationSlack
// @Router  						/settings/projects/{projectId}/integrations/slack/already-set [get]
// @Accept 							json
// @Produce							json
// @Security						Bearer
// @Param        					projectId    	path     	string  				true  	"projectId"
// @Success      					200  			{object} 	model.StateResponse
// @Failure      					400  			{object} 	model.HttpError
// @Failure      					404  			{object} 	model.HttpError
// @Failure      					500  			{object} 	model.HttpError
func (sac *SettingsApiClient) IsSlackCredentialsAlreadySet(c *gin.Context) {
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

	exist, err := svc.IsSlackCredentialsAlreadySet(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to check if slack-credentials exist."),
		})
		return
	}

	c.JSON(http.StatusOK, model.StateResponse{
		State: exist,
	})
}

// CreateMailgunCredentials godoc
// @title           		CreateMailgunCredentials
// @description     		Create the access data for your Mailgun integration.
// @Tags 					IntegrationMailgun
// @Router  				/settings/projects/{projectId}/integrations/mailgun [post]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    				path     	string  						true  	"projectId"
// @Param					MailgunCredentialsRequest 	body 		model.MailgunCredentialsRequest 	true 	"MailgunCredentialsRequest"
// @Success      			200  						{object} 	model.SuccessMessage
// @Failure      			400  						{object} 	model.HttpError
// @Failure      			404  						{object} 	model.HttpError
// @Failure      			500  						{object} 	model.HttpError
func (sac *SettingsApiClient) CreateMailgunCredentials(c *gin.Context) {
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

	var mailgunCredentialRequest model.MailgunCredentialsRequest
	err = c.BindJSON(&mailgunCredentialRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to mailgun-credentials-request-object"})
		return
	}

	host, err := svc.CreateMailgunCredentials(projectId, mailgunCredentialRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to create mailgun-credentials."),
		})
		return
	}

	c.JSON(http.StatusOK, host)
}

// GetMailgunCredentials	godoc
// @title           		GetMailgunCredentials
// @description     		Get the credentials from the mailgun-configuration
// @Tags 					IntegrationMailgun
// @Router  				/settings/projects/{projectId}/integrations/mailgun [get]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    				path     	string  						true  	"projectId"
// @Success      			200  						{object} 	model.SuccessMessage
// @Failure      			400  						{object} 	model.HttpError
// @Failure      			404  						{object} 	model.HttpError
// @Failure      			500  						{object} 	model.HttpError
func (sac *SettingsApiClient) GetMailgunCredentials(c *gin.Context) {
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

	host, err := svc.GetMailgunCredentials(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to get mailgun-credentials."),
		})
		return
	}

	c.JSON(http.StatusOK, host)
}

// UpdateMailgunCredentials godoc
// @title           		UpdateMailgunCredentials
// @description     		Overwrite the token of the Mailgun access of a project
// @Tags 					IntegrationMailgun
// @Router  				/settings/projects/{projectId}/integrations/mailgun [put]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    				path     	string  						true  	"projectId"
// @Param					MailgunCredentialsRequest 	body 		model.MailgunCredentialsRequest 	true 	"MailgunCredentialsRequest"
// @Success      			200  						{object} 	model.SuccessMessage
// @Failure      			400  						{object} 	model.HttpError
// @Failure      			404  						{object} 	model.HttpError
// @Failure      			500  						{object} 	model.HttpError
func (sac *SettingsApiClient) UpdateMailgunCredentials(c *gin.Context) {
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

	var mailgunCredentialRequest model.MailgunCredentialsRequest
	err = c.BindJSON(&mailgunCredentialRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to mailgun-credentials-request-object"})
		return
	}

	host, err := svc.UpdateMailgunCredentials(projectId, mailgunCredentialRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to update mailgun-credentials."),
		})
		return
	}

	c.JSON(http.StatusOK, host)
}

// DeleteMailgunCredentials godoc
// @title           		DeleteMailgunCredentials
// @description     		Deletes mailgun-credentials from a project
// @Tags 					IntegrationMailgun
// @Router  				/settings/projects/{projectId}/integrations/mailgun [delete]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Success      			200  			{object} 	model.SuccessMessage
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) DeleteMailgunCredentials(c *gin.Context) {
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

	successMessage, err := svc.DeleteMailgunCredentials(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to delete mailgun-credentials."),
		})
		return
	}

	c.JSON(http.StatusOK, successMessage)
}

// IsMailgunCredentialsAlreadySet	godoc
// @title           				IsMailgunCredentialsAlreadySet
// @description     				Indicates whether mailgun has already been implemented in a project.
// @Tags 							IntegrationMailgun
// @Router  						/settings/projects/{projectId}/integrations/mailgun/already-set [get]
// @Accept 							json
// @Produce							json
// @Security						Bearer
// @Param        					projectId    	path     	string  				true  	"projectId"
// @Success      					200  			{object} 	model.StateResponse
// @Failure      					400  			{object} 	model.HttpError
// @Failure      					404  			{object} 	model.HttpError
// @Failure      					500  			{object} 	model.HttpError
func (sac *SettingsApiClient) IsMailgunCredentialsAlreadySet(c *gin.Context) {
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

	exist, err := svc.IsMailgunCredentialsAlreadySet(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to check if mailgun-credentials exist."),
		})
		return
	}

	c.JSON(http.StatusOK, model.StateResponse{
		State: exist,
	})
}
