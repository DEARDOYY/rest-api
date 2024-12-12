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

	// ทำ group ที่ต้องการ authen
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authentiate)
	authenticated.POST("/events", crateEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("events/:id/register", registerForEvent)
	authenticated.DELETE("events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("login", login)
}
