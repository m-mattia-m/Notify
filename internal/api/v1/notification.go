package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationApiClient struct{}

type NotificationApi interface {
	SendNotification(c *gin.Context)
}

func NewNotificationApiClient() NotificationApi {
	return &NotificationApiClient{}
}

func (nac *NotificationApiClient) SendNotification(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "works 🚀"})

	//svc, oidcUser, err := getServiceAndUser(c, true)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed"})
	//	return
	//}
	//
	//err = svc.Mailgun.SendMail()
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed"})
	//	return
	//}
	//
	//_ = oidcUser
}
