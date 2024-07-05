// cmd/main.go
package main

import (
	"log"
	"net/http"

	"dev-performance-analytics/pkg/config"
	"dev-performance-analytics/internal/api"
)

func main() {
	// Load the configuration from .env file
	config.LoadConfig()

	// Set up the router
	router := api.SetupRouter()

	// Start the HTTP server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
