package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // handler for inconming http requests (get, post, put, patch, delete)
	server.GET("/events/:id", getEvent) //events/1,  /events/5, etc
	server.POST("/events", createEvent)
}
