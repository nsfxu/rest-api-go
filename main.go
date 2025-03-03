package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", ping)
	server.GET("/events", getEvents)

	server.Run(":8080")
}

func ping(context *gin.Context) {
	context.JSON(http.StatusOK, "pong!")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}
