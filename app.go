package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// This initializes a server
	engine := gin.Default()

	routes.RegisterRoute(engine)

	// This runs the server on this port
	engine.Run(":8080")
}
