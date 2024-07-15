// internal/api/utils.go
package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents the structure for error responses
type ErrorResponse struct {
	Message string `json:"message"`
}

// handleError handles the error by logging it and sending a response to the client
func handleError(c *gin.Context, err error, statusCode int) {
	log.Printf("Error: %v", err)
	c.JSON(statusCode, ErrorResponse{Message: err.Error()})
}
