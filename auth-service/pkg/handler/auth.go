// pkg/handler/auth.go
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/maulikam/auth-service/pkg/model"
	"github.com/maulikam/auth-service/pkg/service"
	"github.com/maulikam/auth-service/pkg/utils"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	UserService service.UserService
}

func NewAuthHandler(userService service.UserService) *AuthHandler {
	return &AuthHandler{UserService: userService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.Authenticate(req.Username, req.Password)
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

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var req signupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
	}

	if err := h.UserService.CreateUser(user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
