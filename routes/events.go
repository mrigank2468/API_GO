package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrigank2468/API_GO/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, events)
}
func createEvents(context *gin.Context) {
	var events models.Event
	err := context.ShouldBindJSON(&events)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		log.Println(err)
		return
	}
	userId := context.GetInt64("userId")
	events.UserId = userId
	err = events.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event"})
		log.Println(err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": events, "status": "success"})

}
func getEventsByID(context *gin.Context) {
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
	context.JSON(http.StatusOK, event)
}
func updateEvents(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		log.Println(err)
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		log.Println(err)
		return
	}
	if event.UserId != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "you are not allowed to update this event"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		log.Println(err)
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully", "event": updatedEvent, "status": "success"})
}

func deleteEvents(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		log.Println(err)
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		log.Println(err)
		return
	}
	if event.UserId != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "you are not allowed to delete this event"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully", "status": "success"})
}
