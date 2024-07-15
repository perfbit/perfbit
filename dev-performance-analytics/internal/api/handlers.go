// internal/api/handlers.go
package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"dev-performance-analytics/pkg/config"
	"dev-performance-analytics/internal/services"
)

// LoginData represents the data required for user login
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler godoc
// @Summary User Login
// @Description User login with username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param loginData body LoginData true "Login Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var loginData LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}

	user, err := config.UserRepository.GetUserByUsername(loginData.Username)
	if err != nil || user.Password != loginData.Password {
		handleError(c, errors.New("invalid username or password"), http.StatusUnauthorized)
		return
	}

	log.Printf("User %s logged in successfully", loginData.Username)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// GetRepositoriesHandler godoc
// @Summary Get Repositories
// @Description Fetch repositories for authenticated user
// @Tags repositories
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} services.Repository
// @Failure 500 {object} common.ErrorResponse
// @Router /repos [get]
func GetRepositoriesHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	log.Println("Fetching repositories")
	repos, err := services.GetRepositories(token)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}

	log.Println("Repositories fetched successfully")
	c.JSON(http.StatusOK, repos)
}

// GetBranchesHandler godoc
// @Summary Get Branches
// @Description Fetch branches for a given repository
// @Tags branches
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Repository owner"
// @Param repo path string true "Repository name"
// @Success 200 {array} services.Branch
// @Failure 500 {object} common.ErrorResponse
// @Router /repos/{id}/branches [get]
func GetBranchesHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	owner := c.Param("id")
	repo := c.Param("repo")

	log.Printf("Fetching branches for owner: %s, repo: %s", owner, repo)
	branches, err := services.GetBranches(token, owner, repo)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}

	log.Println("Branches fetched successfully")
	c.JSON(http.StatusOK, branches)
}

// GetCommitsHandler godoc
// @Summary Get Commits
// @Description Fetch commits for a given branch of a repository
// @Tags commits
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Repository owner"
// @Param repo path string true "Repository name"
// @Param branch path string true "Branch name"
// @Success 200 {array} services.Commit
// @Failure 500 {object} common.ErrorResponse
// @Router /repos/{id}/branches/{branch}/commits [get]
func GetCommitsHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	owner := c.Param("id")
	repo := c.Param("repo")
	branch := c.Param("branch")

	log.Printf("Fetching commits for owner: %s, repo: %s, branch: %s", owner, repo, branch)
	commits, err := services.GetCommits(token, owner, repo, branch)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}

	log.Println("Commits fetched successfully")
	c.JSON(http.StatusOK, commits)
}
