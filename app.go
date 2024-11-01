package main

import (
	"net/http"
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

// Handler is the entry point for Vercel's serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize the database
	db.InitDB()

	// Initialize a Gin engine
	engine := gin.Default()

	// Register routes
	routes.RegisterRoute(engine)

	// Serve the HTTP request
	engine.ServeHTTP(w, r)
}
