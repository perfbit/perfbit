package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"dev-performance-analytics/pkg/config"
	"dev-performance-analytics/internal/services"
)

func LoginHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

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
