// cmd/main.go
package main

import (
	"log"
	"net/http"

	"github.com/yourusername/dev-performance-analytics/pkg/config"
	"github.com/yourusername/dev-performance-analytics/internal/api"
)

func main() {
	config.LoadConfig()

	router := api.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
