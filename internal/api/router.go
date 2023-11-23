package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"notify/internal/api/auth"
	v1 "notify/internal/api/v1"
	"notify/internal/service"
	docs "notify/swagger-docs"
	"strings"
)

func Router(svc service.Service) *gin.Engine {

	if viper.GetString("app.env") == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	swaggerPort := ""
	if viper.GetBool("domain.swagger.port") {
		swaggerPort = fmt.Sprintf(":%s", viper.GetString("server.port"))
	}

	docs.SwaggerInfo.Schemes = []string{viper.GetString("server.scheme")}
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", viper.GetString("server.domain"), swaggerPort)
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s", viper.GetString("server.version"))

	r := gin.Default()
	r.RedirectTrailingSlash = false

	corsPolicy := cors.DefaultConfig()
	corsPolicy.AllowOrigins = []string{"*"}
	corsPolicy.AllowHeaders = append(corsPolicy.AllowHeaders, "Authorization")

	r.Use(cors.New(corsPolicy))
	r.Use(setService(svc))

	r.Use(checkIfRequestFromVerifiedSource(svc))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "👋 OK")
	})
	r.GET("/readiness", func(c *gin.Context) {
		c.JSON(http.StatusOK, "👋 OK")
	})
	r.GET("/liveliness", func(c *gin.Context) {
		// TODO: Check if clients are available and if db available
		c.JSON(http.StatusOK, "👋 OK")
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/v1/swagger/index.html")
	})

	notificationApiClient := v1.NewNotificationApiClient()
	settingsApiClient := v1.NewSettingApiClient()

	v1Group := r.Group("/v1")
	{
		settingsGroup := v1Group.Group("/settings", auth.Authenticate)
		{
			projectGroup := settingsGroup.Group("/projects")
			{
				projectGroup.POST("", settingsApiClient.CreateProject)
				projectGroup.GET("", settingsApiClient.ListProjects)
				projectGroup.GET("/:projectId", settingsApiClient.GetProject)
				projectGroup.PUT("/:projectId", settingsApiClient.UpdateProject)
				projectGroup.DELETE("/:projectId", settingsApiClient.DeleteProject)

				hostGroup := projectGroup.Group("/:projectId/hosts")
				{
					hostGroup.POST("", settingsApiClient.CreateHost)
					hostGroup.GET("", settingsApiClient.ListHosts)
					hostGroup.GET("/:hostId", settingsApiClient.GetHost)
					hostGroup.DELETE("/:hostId", settingsApiClient.DeleteHost)
					hostGroup.PUT("/:hostId/verify", settingsApiClient.VerifyHost)
				}

				integrationGroup := projectGroup.Group("/:projectId/integrations")
				{
					slackGroup := integrationGroup.Group("/slack")
					{
						slackGroup.POST("", settingsApiClient.CreateSlackCredentials)
						slackGroup.GET("/already-set", settingsApiClient.IsSlackCredentialsAlreadySet)
						slackGroup.PUT("", settingsApiClient.UpdateSlackCredentials)
						slackGroup.DELETE("", settingsApiClient.DeleteSlackCredentials)
					}
					mailgunGroup := integrationGroup.Group("/mailgun")
					{
						mailgunGroup.POST("", settingsApiClient.CreateMailgunCredentials)
						mailgunGroup.GET("", settingsApiClient.GetMailgunCredentials)
						mailgunGroup.GET("/already-set", settingsApiClient.IsMailgunCredentialsAlreadySet)
						mailgunGroup.PUT("", settingsApiClient.UpdateMailgunCredentials)
						mailgunGroup.DELETE("", settingsApiClient.DeleteMailgunCredentials)
					}

				}

				flowGroup := projectGroup.Group("/:projectId/flows") // a flow is a notification-workflow which defines
				{
					flowGroup.POST("", settingsApiClient.CreateFlow)           // Create a Flow
					flowGroup.GET("", settingsApiClient.ListFlow)              // List all flows from a project
					flowGroup.GET("/:flowId", settingsApiClient.GetFlow)       // Get a specific flow from a project
					flowGroup.PUT("/:flowId", settingsApiClient.UpdateFlow)    // Update a specific flow from a project
					flowGroup.DELETE("/:flowId", settingsApiClient.DeleteFlow) // Delete a specific flow from a project
				}

				activityGroup := projectGroup.Group("/:projectId/activities")
				{
					activityGroup.GET("", settingsApiClient.ListActivities)
					activityGroup.GET("/:activityId", settingsApiClient.GetActivity)
				}
			}
		}
		notificationGroup := v1Group.Group("/notifications")
		{
			notificationGroup.POST("", notificationApiClient.SendNotification)

		}
		v1Group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}

func setService(svc service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("svc", svc)
		c.Next()
	}
}

func checkIfRequestFromVerifiedSource(svc service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/v1/notification/") {
			return
		}

		verified, err := svc.IfHostVerified(c.Request.Host)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "access denied"})
			return
		}

		if !verified {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "access denied"})
			return
		}
		c.Next()
	}
}
