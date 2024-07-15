package middleware

import (
    "bytes"
    "database/sql"
    "encoding/gob"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/gin-contrib/sessions" 
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        token := session.Get("github_token")
        sessionID := session.ID()
        log.Printf("Session ID at middleware: %s", sessionID)

        if token == nil {
            log.Println("No token found in session, attempting to retrieve from DB")

            // Retrieve token ID from request header or other means
            tokenID := c.GetHeader("Token-ID") // Adjust this as needed to retrieve token ID
            if tokenID == "" {
                log.Println("No token ID provided")
                c.JSON(http.StatusUnauthorized, gin.H{"error": "No token ID provided"})
                c.Abort()
                return
            }

            // Retrieve session information from database
            db, err := sql.Open("postgres", os.Getenv("DATABASE_DSN"))
            if err != nil {
                log.Fatalf("Failed to connect to database: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
                c.Abort()
                return
            }
            defer db.Close()

            // Retrieve session data from DB
            row := db.QueryRow("SELECT data, token FROM http_sessions WHERE token = $1", tokenID)
            var sessionData []byte
            var dbToken string
            if err := row.Scan(&sessionData, &dbToken); err != nil {
                if err == sql.ErrNoRows {
                    log.Println("Session not found in DB")
                    c.JSON(http.StatusUnauthorized, gin.H{"error": "Session not found"})
                } else {
                    log.Printf("Error retrieving session from DB: %v", err)
                    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session from DB"})
                }
                c.Abort()
                return
            }

            // Decode session data
            var sessionMap map[interface{}]interface{}
            decoder := gob.NewDecoder(bytes.NewReader(sessionData))
            if err := decoder.Decode(&sessionMap); err != nil {
                log.Printf("Failed to decode session data: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode session data"})
                c.Abort()
                return
            }

            // Set the session data manually
            session = sessions.DefaultMany(c, "mysession")
            for k, v := range sessionMap {
                session.Set(k, v)
            }

            if err := session.Save(); err != nil {
                log.Printf("Failed to save session data: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session data"})
                c.Abort()
                return
            }

            token = dbToken
        }

        if token == nil {
            log.Println("No token found in session after DB retrieval")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        authHeader := c.GetHeader("Authorization")
        authToken := strings.TrimPrefix(authHeader, "Bearer ")
        if authToken != token.(string) {
            log.Println("Invalid or expired token")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
