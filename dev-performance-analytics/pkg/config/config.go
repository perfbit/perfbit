// pkg/config/config.go
package config

import (
    "dev-performance-analytics/internal/repository"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
    "github.com/joho/godotenv"
)

var (
    DB               *gorm.DB
    UserRepository   repository.UserRepository
    RepoRepository   repository.RepositoryRepository
)

func initDB() {
    dsn := os.Getenv("DATABASE_DSN")
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Database connection established")
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
