package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	db.InitDB()
	defer db.CloseDB()

	server := gin.Default()

	routes.RegusterRoutes(server)

	server.Run(":8080") // localhost:8080
}
