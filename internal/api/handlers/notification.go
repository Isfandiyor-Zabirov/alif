package handlers

import (
	"alif/internal/entities/notifications"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateNotification - Creating notification
// @Summary Creating notification
// @ID create-notification
// @Tags Notifications
// @Produce json
// @Param orderType 	body string true "The type of operation. Can be Purchase, VerifyCard or SendOtp"
// @Param sessionId 	body string true "The session id of the operation"
// @Param card 			body string true "The card pan"
// @Param eventDate 	body string true "The date of the operation"
// @Param websiteUrl 	body string true "The source of operation"
// @Param model body notifications.Event false "Event data"
// @Success 200 {string} string "message"
// @Failure 400 {string} string "reason"
// @Router /api/v1/notifications [post]
func CreateNotification(c *gin.Context) {
	var (
		request notifications.Event
		err     error
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		log.Println("CreateNotification handler cannot bind the request:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Wrong request format"})
		return
	}

	if err = notifications.CreateNotification(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// GetAllNotifications - Getting all notification list
// @Summary Getting all notification list
// @ID get-notifications
// @Tags Notifications
// @Produce json
// @Success 200 {object} []notifications.Notification
// @Failure 400 {string} string "reason"
// @Router /api/v1/notifications [get]
func GetAllNotifications(c *gin.Context) {
	list := notifications.GetAllNotifications()

	if list == nil {
		list = []*notifications.Notification{}
	}

	c.JSON(http.StatusOK, list)
}
