package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // creates web server with basic functionality

	server.GET("/events", getEvents) // handler for inconming http requests (get, post, put, patch, delete)

	server.Run(":8080") // starts this server on localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
