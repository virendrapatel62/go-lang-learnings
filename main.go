package main

import (
	"net/http"

	"github.com/api_app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	error := context.ShouldBindJSON(&event)

	if error == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse the request body",
		})

		return
	}

	event.Id = 1
	event.UserId = 1

	event.Save()

	context.JSON(http.StatusOK, gin.H{
		"message": "event created", "event": event,
	})

}
