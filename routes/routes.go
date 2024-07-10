package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrigank2468/API_GO/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventsByID)
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvents)
	authenticated.DELETE("/events/:id", deleteEvents)
	authenticated.POST("/events/:id/register", register)
	authenticated.DELETE("/events/:id/register", unregister)
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/events/:id/registereduser", registeredUsers)
}
