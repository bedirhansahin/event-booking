package routes

import (
	"example.com/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// events
	authenticated := server.Group("/events")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/", createEvent)
	authenticated.PUT("/:id", updateEvent)
	authenticated.DELETE("/:id", deleteEvent)
	// registration
	authenticated.POST("/:id/register")
	authenticated.DELETE("/:id/register")


	server.POST("/signup", signup)
	server.POST("/login", login)
}
