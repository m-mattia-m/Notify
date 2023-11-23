package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify/internal/helper"
	"notify/internal/model"
)

type NotificationApiClient struct{}

type NotificationApi interface {
	SendNotification(c *gin.Context)
}

func NewNotificationApiClient() NotificationApi {
	return &NotificationApiClient{}
}

// SendNotification 		godoc
// @title           		SendNotification
// @description     		Send a notification
// @Tags 					Notification
// @Router  				/notifications [post]
// @Accept 					json
// @Produce					json
// @Param					Notification 		body 		model.Notification 		true 	"Notification"
// @Success      			200  				{object} 	model.SuccessMessage
// @Failure      			400  				{object} 	model.HttpError
// @Failure      			404  				{object} 	model.HttpError
// @Failure      			500  				{object} 	model.HttpError
func (nac *NotificationApiClient) SendNotification(c *gin.Context) {
	svc, _, err := getServiceAndUser(c, false)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed"})
		return
	}

	var notification model.Notification
	err = c.BindJSON(&notification)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to convert/bind body to notification-object"})
		return
	}

	host := c.Request.Host
	successMessage, err := svc.SendNotification(host, notification)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": helper.ValidateErrorResponse(err, "failed to send"),
		})
		return
	}

	c.JSON(http.StatusOK, successMessage)
}
