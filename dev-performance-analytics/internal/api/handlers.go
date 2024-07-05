package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "dev-performance-analytics/internal/services"
    "dev-performance-analytics/pkg/config"
)

var users = map[string]string{
    "user1": "password1",
    "user2": "password2",
}

func loginHandler(c *gin.Context) {
    var loginData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    expectedPassword, exists := users[loginData.Username]
    if !exists || expectedPassword != loginData.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func getRepositoriesHandler(c *gin.Context) {
    token := c.GetHeader("Authorization")
    repos, err := services.GetRepositories(token)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, repos)
}

func getBranchesHandler(c *gin.Context) {
    token := c.GetHeader("Authorization")
    owner := c.Param("id")
    repo := c.Param("repo")

    branches, err := services.GetBranches(token, owner, repo)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, branches)
}

func getCommitsHandler(c *gin.Context) {
    token := c.GetHeader("Authorization")
    owner := c.Param("id")
    repo := c.Param("repo")
    branch := c.Param("branch")

    commits, err := services.GetCommits(token, owner, repo, branch)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, commits)
}
