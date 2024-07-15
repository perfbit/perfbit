// pkg/config/config.go
package config

import (
    "dev-performance-analytics/internal/repository"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    DB               *gorm.DB
    UserRepository   repository.UserRepository
    RepoRepository   repository.RepositoryRepository
)

func initDB() {
    dsn := os.Getenv("DATABASE_DSN")
    var err error

    // Retry logic
    maxRetries := 3
    for i := 0; i < maxRetries; i++ {
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            log.Println("Database connection established")
            return
        }
        log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
        time.Sleep(5 * time.Second)
    }

    // If we reach here, all retries failed
    log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
}

func initRepositories() {
    UserRepository = repository.NewUserRepository(DB)
    RepoRepository = repository.NewRepositoryRepository(DB)
}

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Initialize database
    initDB()
    // Initialize repositories
    initRepositories()
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
