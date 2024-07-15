// internal/api/router.go
package api

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "github.com/google/go-github/v63/github"
    "golang.org/x/oauth2"
    gh "golang.org/x/oauth2/github"
    "dev-performance-analytics/common"
    "dev-performance-analytics/pkg/middleware"
)

var (
    githubOAuthConfig *oauth2.Config
    oauthStateString  = "random"
)

func init() {
    githubOAuthConfig = &oauth2.Config{
        RedirectURL:  "http://localhost:8080/auth/github/callback",
        ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
        ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
        Scopes:       []string{"repo", "user"},
        Endpoint:     gh.Endpoint,
    }
}

// SetupRouter initializes the Gin router with all the endpoints
func SetupRouter(router *gin.Engine) {
    // Configure CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

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
        v1.GET("/dashboard", getDashboardData)
    }
}

// handleGitHubLogin godoc
// @Summary GitHub Login
// @Description Redirect to GitHub login
// @Tags auth
// @Produce json
// @Success 302
// @Router /auth/github/login [get]
func handleGitHubLogin(c *gin.Context) {
    log.Println("Redirecting to GitHub login")
    url := githubOAuthConfig.AuthCodeURL(oauthStateString)
    c.Redirect(http.StatusTemporaryRedirect, url)
}

// handleGitHubCallback godoc
// @Summary GitHub Callback
// @Description Handle GitHub callback and authenticate user
// @Tags auth
// @Produce json
// @Success 302
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /auth/github/callback [get]
func handleGitHubCallback(c *gin.Context) {
    state := c.Query("state")
    if state != oauthStateString {
        log.Println("Invalid state")
        c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid state"})
        return
    }

    code := c.Query("code")
    log.Printf("Exchanging code: %s", code)
    token, err := githubOAuthConfig.Exchange(context.Background(), code)
    if err != nil {
        log.Printf("Failed to exchange token: %v", err)
        c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "failed to exchange token"})
        return
    }

    client := github.NewClient(githubOAuthConfig.Client(context.Background(), token))
    user, _, err := client.Users.Get(context.Background(), "")
    if err != nil {
        log.Printf("Failed to get user: %v", err)
        c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "failed to get user"})
        return
    }

    log.Printf("User %s authenticated successfully", user.GetLogin())

    // Print and log the token
    log.Printf("Token: %s", token.AccessToken)
    fmt.Printf("Token: %s\n", token.AccessToken)

    // Save the OAuth token and user information in the session
    session := sessions.Default(c)
    session.Set("github_token", token.AccessToken)
    session.Set("github_user", user.GetLogin())
    if err := session.Save(); err != nil {
        log.Printf("Failed to save session: %v", err)
        c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "failed to save session"})
        return
    }

    // Log session ID for debugging
    sessionID := session.ID()
    if sessionID == "" {
        log.Println("Session ID is empty")
    } else {
        log.Printf("Session ID: %s", sessionID)
    }

    // Redirect to the frontend with the session token
    redirectURL := fmt.Sprintf("http://localhost:3000/login?token=%s", token.AccessToken)
    c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
