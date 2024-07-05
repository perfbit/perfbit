// internal/api/dashboard.go
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maulikam/dev-performance-analytics/internal/services"
)

func getDashboardData(c *gin.Context) {
	token := c.GetHeader("Authorization")
	owner := c.Query("owner")
	repo := c.Query("repo")
	branch := c.Query("branch")

	commits, err := services.GetCommits(token, owner, repo, branch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := services.GeneratePerformanceMetrics(commits)
	c.JSON(http.StatusOK, data)
}
