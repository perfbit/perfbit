// internal/api/router.go
package api

import (
	"github.com/gin-gonic/gin"
	"dev-performance-analytics/pkg/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
	    // internal/api/router.go
        v1.Use(middleware.AuthMiddleware())
		v1.POST("/login", LoginHandler)
		v1.GET("/repos", GetRepositoriesHandler)
		v1.GET("/repos/:id/branches", GetBranchesHandler)
		v1.GET("/repos/:id/branches/:branch/commits", GetCommitsHandler)
	}

	return router
}

// Add handler functions here
