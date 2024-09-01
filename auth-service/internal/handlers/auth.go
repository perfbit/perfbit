package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/perfbit/perfbit/auth-service/internal/config"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	config       *config.Config
	provider     *oidc.Provider
	oauth2Config oauth2.Config
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

	return &AuthHandler{
		config:       cfg,
		provider:     provider,
		oauth2Config: oauth2Config,
	}, nil
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, h.oauth2Config.AuthCodeURL("state"), http.StatusFound)
}

func (h *AuthHandler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	// Implement OIDC callback handling
}

func (h *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	// Implement logout
}

func (h *AuthHandler) HandleUserInfo(w http.ResponseWriter, r *http.Request) {
	// Implement user info retrieval
}