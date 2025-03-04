package main

import (
	"net/http"

	"example.com/rest-api-go/db"
	"example.com/rest-api-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.GET("/ping", ping)
	server.Run(":8080")
}

func ping(context *gin.Context) {
	context.JSON(http.StatusOK, "pong!")
}
