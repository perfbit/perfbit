package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    cors "github.com/gin-contrib/cors"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "github.com/google/go-github/v39/github"
    "github.com/joho/godotenv"
    "golang.org/x/oauth2"
    gh "golang.org/x/oauth2/github"

    "dev-performance-analytics/internal/api"
    "dev-performance-analytics/internal/models"
    "dev-performance-analytics/pkg/config"
    "dev-performance-analytics/pkg/middleware"

    _ "dev-performance-analytics/docs" // This line is necessary for go-swagger to find your docs!

    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
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

    // Log environment variables
    log.Println("GITHUB_CLIENT_ID:", os.Getenv("GITHUB_CLIENT_ID"))
    log.Println("GITHUB_CLIENT_SECRET:", os.Getenv("GITHUB_CLIENT_SECRET"))
    log.Println("DATABASE_DSN:", os.Getenv("DATABASE_DSN"))

    githubOAuthConfig = &oauth2.Config{
        RedirectURL:  "http://localhost:8080/auth/github/callback",
        ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
        ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
        Scopes:       []string{"repo", "user"},
        Endpoint:     gh.Endpoint,
    }
}

// @title Developer Performance Analytics API
// @version 1.0
// @description This is a developer performance analytics server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
    config.LoadConfig()

    // Migrate the schema
    err := config.DB.AutoMigrate(&models.User{}, &models.Repository{}, &models.Branch{}, &models.Commit{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    router := gin.Default()

    // Session middleware
    store := cookie.NewStore([]byte("secret"))
    router.Use(sessions.Sessions("mysession", store))

    // CORS configuration
    router.Use(cors.New(cors.Config{
        AllowWildcard:       true,
        AllowOrigins:        []string{"http://localhost:3000"},
        AllowMethods:        []string{"PUT", "GET", "POST", "DELETE"},
        AllowHeaders:        []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:       []string{},
        AllowCredentials:    true,
        MaxAge:              50 * time.Second,
        AllowPrivateNetwork: true,
    }))

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

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

    // Save the OAuth token and user information in the session
    session := sessions.Default(c)
    session.Set("github_token", token.AccessToken)
    session.Set("github_user", user.GetLogin())
    session.Save()

    // Redirect to the frontend with the session token
    redirectURL := fmt.Sprintf("http://localhost:3000/login?token=%s", token.AccessToken)
    c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
