package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrigank2468/API_GO/models"
)

func register(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		log.Println(err)
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		log.Println(err)
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for event"})
		log.Println(err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "registered for event successfully", "event": event, "status": "success"})
}
func unregister(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		log.Println(err)
		return
	}
	event, _ := models.GetEventByID(eventId)
	event.ID = eventId
	err=event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not unregister for event"})
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "unregistered for event successfully", "event": event, "status": "success"})
}
func registeredUsers(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		log.Println(err)
		return
	}
	users, err := models.GetRegisteredUsers(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch registered users"})
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, users)
}
