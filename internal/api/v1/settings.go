package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"message-proxy/internal/helper"
	"message-proxy/internal/model"
	"net/http"
)

type SettingsApiClient struct{}

type SettingsApi interface {
	CreateProject(c *gin.Context)
	GetProject(c *gin.Context)
	ListProjects(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)

	CreateHost(c *gin.Context)
	ListHosts(c *gin.Context)
	GetHost(c *gin.Context)
	VerifyHost(c *gin.Context)
	DeleteHost(c *gin.Context)
}

func NewSettingApiClient() SettingsApi {
	return &SettingsApiClient{}
}

// CreateProject 			godoc
// @title           		CreateProject
// @description     		Creates a project from the request body which is also sent and return this project.
// @Tags 					Project
// @Router  				/settings/projects [post]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param					ProjectRequest 	body 		model.ProjectRequest 	true 	"ProjectRequest"
// @Success      			200  			{object} 	model.Project
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) CreateProject(c *gin.Context) {
	svc, oidcUser, err := getServiceAndUser(c, true)
	if err != nil {
		log.Error("failed to get service or user from context :: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "an unexpected error has occurred"})
		return
	}

	var projectRequest model.ProjectRequest
	err = c.BindJSON(&projectRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to project-request-object"})
		return
	}

	project, err := svc.CreateProject(projectRequest, oidcUser.Sub)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to create project."),
		})
		return
	}

	c.JSON(http.StatusOK, project)
}

// ListProjects 			godoc
// @title           		ListProjects
// @description     		List all projects from the signed-in user
// @Tags 					Project
// @Router  				/settings/projects [get]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Success      			200  			{object} 	[]model.Project
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) ListProjects(c *gin.Context) {
	svc, oidcUser, err := getServiceAndUser(c, true)
	if err != nil {
		log.Error("failed to get service or user from context :: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "an unexpected error has occurred"})
		return
	}

	projects, err := svc.ListProjects(oidcUser.Sub)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to list projects."),
		})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// GetProject 				godoc
// @title           		GetProject
// @description     		Returns a project by their id.
// @Tags 					Project
// @Router  				/settings/projects/{projectId} [get]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Success      			200  			{object} 	model.Project
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) GetProject(c *gin.Context) {
	svc, oidcUser, err := getServiceAndUser(c, true)
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

	project, err := svc.GetProject(projectId, oidcUser.Sub)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to get project."),
		})
		return
	}

	c.JSON(http.StatusOK, project)
}

// UpdateProject 			godoc
// @title           		UpdateProject
// @description     		Updates a project by their id from the request body which is also sent and return this project.
// @Tags 					Project
// @Router  				/settings/projects/{projectId} [put]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Param					ProjectRequest 	body 		model.ProjectRequest 	true 	"ProjectRequest"
// @Success      			200  			{object} 	model.Project
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) UpdateProject(c *gin.Context) {
	svc, oidcUser, err := getServiceAndUser(c, true)
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

	var projectRequest model.ProjectRequest
	err = c.BindJSON(&projectRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to project-request-object"})
		return
	}

	project, err := svc.UpdateProject(projectId, projectRequest, oidcUser.Sub)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to update project."),
		})
		return
	}

	c.JSON(http.StatusOK, project)
}

// DeleteProject 			godoc
// @title           		DeleteProject
// @description     		Deletes a project by their id.
// @Tags 					Project
// @Router  				/settings/projects/{projectId} [delete]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Success      			200  			{object} 	model.SuccessMessage
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) DeleteProject(c *gin.Context) {
	svc, oidcUser, err := getServiceAndUser(c, true)
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

	successMessage, err := svc.DeleteProject(projectId, oidcUser.Sub)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to delete project."),
		})
		return
	}

	c.JSON(http.StatusOK, successMessage)
}

// CreateHost 				godoc
// @title           		CreateHost
// @description     		Creates a host from the request body and the projectId which is also sent and return this host.
// @Tags 					Host
// @Router  				/settings/projects/{projectId}/hosts [post]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Param					HostRequest 	body 		model.HostRequest 	true 	"HostRequest"
// @Success      			200  			{object} 	model.Host
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) CreateHost(c *gin.Context) {
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

	var hostRequest model.HostRequest
	err = c.BindJSON(&hostRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to host-request-object"})
		return
	}

	host, err := svc.CreateHost(hostRequest, projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to create host."),
		})
		return
	}

	c.JSON(http.StatusOK, host)
}

// ListHosts 				godoc
// @title           		ListHosts
// @description     		List all hosts from the given project
// @Tags 					Host
// @Router  				/settings/projects/{projectId}/hosts [get]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Success      			200  			{object} 	[]model.Host
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) ListHosts(c *gin.Context) {
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

	hosts, err := svc.ListHosts(projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to list hosts."),
		})
		return
	}

	c.JSON(http.StatusOK, hosts)
}

// GetHost 					godoc
// @title           		GetHost
// @description     		Returns a host by their id and their projectId.
// @Tags 					Host
// @Router  				/settings/projects/{projectId}/hosts/{hostId} [get]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Param        			hostId    		path     	string  				true  	"hostId"
// @Success      			200  			{object} 	model.Host
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) GetHost(c *gin.Context) {
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

	hostId := c.Param("hostId")
	if hostId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "hostId is required"})
		return
	}

	host, err := svc.GetHost(hostId, projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to get host."),
		})
		return
	}

	c.JSON(http.StatusOK, host)
}

// VerifyHost 				godoc
// @title           		VerifyHost
// @description     		trigger a verification from the given host
// @Tags 					Host
// @Router  				/settings/projects/{projectId}/hosts/{hostId}/verify [put]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Param        			hostId    		path     	string  				true  	"hostId"
// @Success      			200  			{object} 	model.Host
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) VerifyHost(c *gin.Context) {
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

	hostId := c.Param("hostId")
	if hostId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "hostId is required"})
		return
	}

	host, err := svc.VerifyHost(hostId, projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to update host."),
		})
		return
	}

	c.JSON(http.StatusOK, host)
}

// DeleteHost 				godoc
// @title           		DeleteHost
// @description     		Deletes a host by their id and their projectId.
// @Tags 					Host
// @Router  				/settings/projects/{projectId}/hosts/{hostId} [delete]
// @Accept 					json
// @Produce					json
// @Security				Bearer
// @Param        			projectId    	path     	string  				true  	"projectId"
// @Param        			hostId    		path     	string  				true  	"hostId"
// @Success      			200  			{object} 	model.SuccessMessage
// @Failure      			400  			{object} 	model.HttpError
// @Failure      			404  			{object} 	model.HttpError
// @Failure      			500  			{object} 	model.HttpError
func (sac *SettingsApiClient) DeleteHost(c *gin.Context) {
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

	hostId := c.Param("hostId")
	if hostId == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "hostId is required"})
		return
	}

	successMessage, err := svc.DeleteHost(hostId, projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to delete host."),
		})
		return
	}

	c.JSON(http.StatusOK, successMessage)
}
