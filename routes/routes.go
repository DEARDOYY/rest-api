package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegusterRoutes(server *gin.Engine) {
	// GET ,POST ,PUT ,PATCH ,DELETE
	// ถ้าอยู่ใน package เดียวกันไม่ต้อง import
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.Authentiate, crateEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("login", login)
}
