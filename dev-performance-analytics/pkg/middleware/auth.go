// pkg/middleware/auth.go
package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        token := session.Get("github_token")
        if token == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        authHeader := c.GetHeader("Authorization")
        authToken := strings.TrimPrefix(authHeader, "Bearer ")
        if authToken != token.(string) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }
        c.Next()
    }
}
