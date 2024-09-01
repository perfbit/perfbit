package handlers

import (
	"context"
	"encoding/json"
	"net/http"
    "time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/uuid"
	"github.com/perfbit/perfbit/auth-service/internal/config"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	config       *config.Config
	provider     *oidc.Provider
	oauth2Config oauth2.Config
	redisClient  *storage.RedisClient
}

func (h *AuthHandler) RedisClient() *storage.RedisClient {
    return h.redisClient
}

func NewAuthHandler(cfg *config.Config) (*AuthHandler, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, cfg.AuthentikURL)
	if err != nil {
		return nil, err
	}

	oauth2Config := oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	redisClient, err := storage.NewRedisClient(cfg)
    if err != nil {
        return nil, err
    }

    return &AuthHandler{
        config:       cfg,
        provider:     provider,
        oauth2Config: oauth2Config,
        redisClient:  redisClient,
    }, nil
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, h.oauth2Config.AuthCodeURL("state"), http.StatusFound)
}

func (h *AuthHandler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	oauth2Token, err := h.oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	idToken, err := h.provider.Verifier(&oidc.Config{ClientID: h.config.ClientID}).Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var claims struct {
		Email    string `json:"email"`
		Username string `json:"preferred_username"`
	}
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a session ID
	sessionID := uuid.New().String()

	// Store session in Redis
	sessionData := map[string]interface{}{
		"email":    claims.Email,
		"username": claims.Username,
		"token":    rawIDToken,
	}
	err = h.redisClient.Set(r.Context(), sessionID, sessionData, 24*time.Hour)
	if err != nil {
		http.Error(w, "Failed to create session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(24 * time.Hour.Seconds()),
	})

	// Redirect to a welcome page or dashboard
	http.Redirect(w, r, "/welcome", http.StatusFound)
}

func (h *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, you would invalidate the user's session here
	// For now, we'll just redirect to Authentik's logout endpoint
	logoutURL := h.config.AuthentikURL + "/application/o/logout/"
	http.Redirect(w, r, logoutURL, http.StatusFound)
}

func (h *AuthHandler) HandleUserInfo(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    // Get the token from the context (set by the middleware)
    idToken, ok := ctx.Value("token").(*oidc.IDToken)
    if !ok {
        http.Error(w, "No valid token found", http.StatusUnauthorized)
        return
    }

    var claims map[string]interface{}
    if err := idToken.Claims(&claims); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(claims)
}

func (h *AuthHandler) Provider() *oidc.Provider {
    return h.provider
}