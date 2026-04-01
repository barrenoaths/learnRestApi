package main

import (
	"learnRestApi/db"
	"learnRestApi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() // creates web server with basic functionality
	routes.RegisterRoutes(server)
	server.Run(":8080") // starts this server on localhost:8080
}
