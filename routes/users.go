package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrigank2468/API_GO/models"
	"github.com/mrigank2468/API_GO/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = user.Save()
	if err != nil {	
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}
	context.JSON(http.StatusCreated,gin.H{"message":"user created succesfully"})	
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}
	token,err:=utils.GenerateToken(user.Email,user.ID)
	if err!= nil{ 
		context.JSON(http.StatusInternalServerError,gin.H{"message":"could not validate user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}