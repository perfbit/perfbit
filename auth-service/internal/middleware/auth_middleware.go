package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/perfbit/perfbit/auth-service/internal/config"
)

type AuthMiddleware struct {
    config      *config.Config
    verifier    *oidc.IDTokenVerifier
    redisClient *storage.RedisClient
}

func NewAuthMiddleware(cfg *config.Config, provider *oidc.Provider, redisClient *storage.RedisClient) *AuthMiddleware {
    return &AuthMiddleware{
        config:      cfg,
        verifier:    provider.Verifier(&oidc.Config{ClientID: cfg.ClientID}),
        redisClient: redisClient,
    }
}

func (m *AuthMiddleware) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "No session found", http.StatusUnauthorized)
			return
		}

		var sessionData map[string]interface{}
		err = m.redisClient.Get(r.Context(), cookie.Value, &sessionData)
		if err != nil {
			http.Error(w, "Invalid session", http.StatusUnauthorized)
			return
		}

		token, ok := sessionData["token"].(string)
		if !ok {
			http.Error(w, "No token in session", http.StatusUnauthorized)
			return
		}

		idToken, err := m.verifier.Verify(r.Context(), token)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Add the token to the request context
		ctx := context.WithValue(r.Context(), "token", idToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}