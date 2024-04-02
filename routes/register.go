package routes

import (
	"strconv"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(500, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(500, gin.H{"message": "Could not register user for event"})
		return
	}

	context.JSON(201, gin.H{"message": "Registered!"})

}

func cancelRegistration(context *gin.Context) {

}
