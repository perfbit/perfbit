// pkg/handler/auth.go
package handler

import (
	"encoding/json"
	"github.com/maulikam/auth-service/pkg/service"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req loginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user, err := userService.Authenticate(req.Username, req.Password)
		if err != nil {
			http.Error(w, "Error authenticating user", http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		w.Write([]byte("Login successful"))
	}
}
