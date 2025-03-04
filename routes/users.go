package routes

import (
	"net/http"

	"example.com/rest-api-go/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.Users

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created!"})

}
