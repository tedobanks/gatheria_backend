package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Initialize Gin
	engine := gin.Default()

	// Register routes
	routes.RegisterRoute(engine)

	// Get PORT from environment variable (Render provides this dynamically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if not set
	}

	// Start the server
	fmt.Println("Server running on port:", port)
	engine.Run(":" + port)
}
