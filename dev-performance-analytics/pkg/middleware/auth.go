// pkg/middleware/auth.go
package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"dev-performance-analytics/pkg/config"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token != config.GetEnv("GITHUB_TOKEN") {  // Validate against expected token
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        c.Next()
    }
}
