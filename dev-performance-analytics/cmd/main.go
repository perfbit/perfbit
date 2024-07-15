package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "dev-performance-analytics/internal/api"
    "dev-performance-analytics/internal/models"
    "dev-performance-analytics/pkg/config"

    _ "dev-performance-analytics/docs" // for go-swagger to find docs!

    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
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

    router := api.SetupRouter()

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
