// internal/api/dashboard.go
package api

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"dev-performance-analytics/internal/services"
)

// getDashboardData godoc
// @Summary Get Dashboard Data
// @Description Fetch performance metrics for a repository
// @Tags dashboard
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param owner query string true "Repository owner"
// @Param repo query string true "Repository name"
// @Param branch query string true "Branch name"
// @Success 200 {object} services.PerformanceMetrics
// @Failure 500 {object} ErrorResponse
// @Router /dashboard [get]
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
