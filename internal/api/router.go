package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"message-proxy/docs"
	v1 "message-proxy/internal/api/v1"
	"message-proxy/internal/service"
	"net/http"
)

func Router(svc service.Service) *gin.Engine {

	if viper.GetString("app.env") == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	docs.SwaggerInfo.Schemes = []string{viper.GetString("server.scheme")}
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", viper.GetString("server.domain"), viper.GetString("server.port"))
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s", viper.GetString("server.version"))

	r := gin.Default()
	r.RedirectTrailingSlash = false

	corsPolicy := cors.DefaultConfig()
	corsPolicy.AllowOrigins = []string{
		viper.GetString("frontend.url"),
		fmt.Sprintf("%s://%s:%s", viper.GetString("server.scheme"), viper.GetString("server.domain"), viper.GetString("server.port")),
	}
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

	notificationApiClient := v1.NewNotificationApiClient()

	v1Group := r.Group("/v1")
	{
		settingsGroup := v1Group.Group("/settings")
		{
			projectGroup := settingsGroup.Group("/projects")
			{
				projectGroup.POST("")              // Create Project
				projectGroup.GET("")               // List Projects
				projectGroup.GET("/:projectId")    // Get Project
				projectGroup.PUT("/:projectId")    // Update Project
				projectGroup.DELETE("/:projectId") // Delete Project

				hostGroup := projectGroup.Group("/:projectId/hosts")
				{
					hostGroup.POST("")           // add host
					hostGroup.GET("")            // list hosts
					hostGroup.GET("/:hostId")    // get host
					hostGroup.PUT("/:hostId")    // update host
					hostGroup.DELETE("/:hostId") // delete host
					hostGroup.PUT("/:hostId/verify")
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

// setService: TODO: if you get here a error or one with the svc, change the param from *service.Client (struct) to *service.Service (interface)
func setService(svc service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("svc", svc)
		c.Next()
	}
}

func checkIfRequestFromVerifiedSource(svc service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		clientHost := c.Request.Host

		verified, err := svc.IfHostOrIpVerified(clientIP, clientHost)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "access denied"})
			return
		}

		if verified {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "access denied"})
			return
		}
		c.Next()
	}
}
