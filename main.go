package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrigank2468/API_GO/db"
	"github.com/mrigank2468/API_GO/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
    routes.RegisterRoutes(server)
	server.Run(":8080")
}
