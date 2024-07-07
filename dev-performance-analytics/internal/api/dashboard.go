package api

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"dev-performance-analytics/internal/services"
)

func handleError(c *gin.Context, err error, statusCode int) {
	log.Printf("Error: %v", err)
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func getDashboardData(c *gin.Context) {
	token := c.GetHeader("Authorization")
	owner := c.Query("owner")
	repo := c.Query("repo")
	branch := c.Query("branch")

	log.Printf("Fetching commits for owner: %s, repo: %s, branch: %s", owner, repo, branch)
	commits, err := services.GetCommits(token, owner, repo, branch)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}

	log.Println("Generating performance metrics")
	data := services.GeneratePerformanceMetrics(commits)
	log.Println("Performance metrics generated successfully")
	c.JSON(http.StatusOK, data)
}
