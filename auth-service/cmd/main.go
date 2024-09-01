package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/perfbit/perfbit/auth-service/internal/config"
	"github.com/perfbit/perfbit/auth-service/internal/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := mux.NewRouter()
	
	authHandler, err := handlers.NewAuthHandler(cfg)
	if err != nil {
		log.Fatalf("Failed to create auth handler: %v", err)
	}

	r.HandleFunc("/login", authHandler.HandleLogin)
	r.HandleFunc("/callback", authHandler.HandleCallback)
	r.HandleFunc("/logout", authHandler.HandleLogout)
	r.HandleFunc("/userinfo", authHandler.HandleUserInfo)

	log.Printf("Starting auth service on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}