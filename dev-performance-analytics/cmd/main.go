// cmd/main.go
package main

import (
	"log"
	"net/http"

	"dev-performance-analytics/pkg/config"
	"dev-performance-analytics/internal/api"
)

func main() {
	config.LoadConfig()

	// Retrieve the GitHub token from the environment
	githubToken := config.GetEnv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatal("GITHUB_TOKEN is not set in the environment")
	}

	router := api.SetupRouter()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
