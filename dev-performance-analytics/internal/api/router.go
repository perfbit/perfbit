// internal/api/router.go
package api

import (
	"github.com/gin-gonic/gin"
	"dev-performance-analytics/pkg/middleware"
)

// SetupRouter initializes the Gin router with all the endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Add GitHub OAuth routes
	router.GET("/auth/github/login", handleGitHubLogin)
	router.GET("/auth/github/callback", handleGitHubCallback)

	v1 := router.Group("/api/v1")
	{
		v1.Use(middleware.AuthMiddleware())
		v1.POST("/login", LoginHandler)
		v1.GET("/repos", GetRepositoriesHandler)
		v1.GET("/repos/:id/branches", GetBranchesHandler)
		v1.GET("/repos/:id/branches/:branch/commits", GetCommitsHandler)
		v1.GET("/dashboard", getDashboardData) // Added this line to register the dashboard endpoint
	}

	return router
}
