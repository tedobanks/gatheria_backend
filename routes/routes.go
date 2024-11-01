package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {

	// This is one way to define a MiddleWare, since we have more than one request using our middleware we use server.Group() and we pass in something all the parts that would use that Middleware have in their paths/routes. In our case we pass in / and we just use our variable to make a request as normal.
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenicate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("events/:id/register", registerForEvent)

	// Gets all events
	server.GET("/events", getEvents)

	// Gets a single event
	server.GET("/events/:id", getEvent)

	// DELETE all events
	// This is another way to add middleware to your request. You just have to add the middleware before your request handler / handlers
	server.DELETE("/events/empty", middlewares.Authenicate, emptyEvents)

	// Create a new user
	server.POST("/register", signup)

	// GET a single user
	server.GET("/users/:id", getUser)

	// GET all users
	server.GET("/users", getAllUsers)

	server.POST("/login", login)

}
