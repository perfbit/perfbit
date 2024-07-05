package main

import (
    "context"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/google/go-github/v39/github"
    "github.com/joho/godotenv"
    "golang.org/x/oauth2"
    gh "golang.org/x/oauth2/github"

    "dev-performance-analytics/internal/api"
    "dev-performance-analytics/pkg/config"
    "dev-performance-analytics/pkg/middleware"
)

var (
    githubOAuthConfig *oauth2.Config
    oauthStateString  = "random"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    githubOAuthConfig = &oauth2.Config{
        RedirectURL:  "http://localhost:3000/auth/github/callback",
        ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
        ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
        Scopes:       []string{"user:email"},
        Endpoint:     gh.Endpoint,
    }
}

func main() {
    config.LoadConfig()

    router := gin.Default()

    // Set up routes
    router.GET("/auth/github/login", handleGitHubLogin)
    router.GET("/auth/github/callback", handleGitHubCallback)

    apiGroup := router.Group("/api")
    {
        v1 := apiGroup.Group("/v1")
        {
            v1.Use(middleware.AuthMiddleware()) // Your existing middleware if any
            v1.POST("/login", api.LoginHandler)
            v1.GET("/repos", api.GetRepositoriesHandler)
            v1.GET("/repos/:id/branches", api.GetBranchesHandler)
            v1.GET("/repos/:id/branches/:branch/commits", api.GetCommitsHandler)
        }
    }

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func handleGitHubLogin(c *gin.Context) {
    url := githubOAuthConfig.AuthCodeURL(oauthStateString)
    c.Redirect(http.StatusTemporaryRedirect, url)
}

func handleGitHubCallback(c *gin.Context) {
    state := c.Query("state")
    if state != oauthStateString {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid state"})
        return
    }

    code := c.Query("code")
    token, err := githubOAuthConfig.Exchange(context.Background(), code)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange token"})
        return
    }

    client := github.NewClient(githubOAuthConfig.Client(context.Background(), token))
    user, _, err := client.Users.Get(context.Background(), "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token.AccessToken, "user": user})
}
