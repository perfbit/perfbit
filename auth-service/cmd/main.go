// cmd/main.go
package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/maulikam/auth-service/internal/config"
	"github.com/maulikam/auth-service/pkg/handler"
	"github.com/maulikam/auth-service/pkg/repository"
	"github.com/maulikam/auth-service/pkg/service"
	"github.com/pressly/goose/v3"
)

func main() {
	connStr := "host=" + config.GetEnv("POSTGRES_HOST", "localhost") +
		" port=" + config.GetEnv("POSTGRES_PORT", "5432") +
		" user=" + config.GetEnv("POSTGRES_USER", "postgres") +
		" password=" + config.GetEnv("POSTGRES_PASSWORD", "secret") +
		" dbname=" + config.GetEnv("POSTGRES_DB", "authdb") +
		" sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Run migrations
	if err := goose.Up(db, "db/migrations"); err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewPostgresUserRepository(db)
	userService := service.UserService{Repo: userRepo}
	authHandler := handler.NewAuthHandler(userService)

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/signup", authHandler.Signup)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
