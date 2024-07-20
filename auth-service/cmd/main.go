// cmd/main.go
package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/cors"
	"log"
	"net/http"

	"github.com/maulikam/perfbit/auth-service/internal/config"
	"github.com/maulikam/perfbit/auth-service/pkg/handler"
	"github.com/maulikam/perfbit/auth-service/pkg/middleware"
	"github.com/maulikam/perfbit/auth-service/pkg/repository"
	"github.com/maulikam/perfbit/auth-service/pkg/service"
)

func main() {
	connStr := "host=" + config.GetEnv("POSTGRES_HOST", "localhost") +
		" port=" + config.GetEnv("POSTGRES_PORT", "5433") +
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

	mux := http.NewServeMux()
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/signup", authHandler.Signup)
	mux.HandleFunc("/verify", authHandler.Verify)
	mux.HandleFunc("/refresh", authHandler.Refresh)
	mux.HandleFunc("/auth/github", authHandler.HandleGitHubLogin)
	mux.HandleFunc("/callback", authHandler.HandleGitHubCallback)

	// Protected routes
	protected := http.NewServeMux()
	protected.Handle("/protected-endpoint", middleware.JWTAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a protected endpoint"))
	})))

	mux.Handle("/protected-endpoint", protected)

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handlers := c.Handler(mux)

	http.Handle("/", handlers) // Handle root to handlers
	log.Println("Server started at :8081")
	log.Fatal(http.ListenAndServe(":8081", handlers))
}
