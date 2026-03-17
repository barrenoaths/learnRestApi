package main

import (
	"net/http"

	"learnRestApi/db"
	"learnRestApi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() // creates web server with basic functionality

	server.GET("/events", getEvents) // handler for inconming http requests (get, post, put, patch, delete)
	server.POST("/events", createEvent)
	server.GET("/events2", getEvents) // handler for inconming http requests (get, post, put, patch, delete)
	server.POST("/events2", createEvent)

	server.Run(":8080") // starts this server on localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
