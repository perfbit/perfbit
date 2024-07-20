package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/maulikam/perfbit/auth-service/pkg/email"
	"github.com/maulikam/perfbit/auth-service/pkg/model"
	"github.com/maulikam/perfbit/auth-service/pkg/service"
	"github.com/maulikam/perfbit/auth-service/pkg/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"
)

var (
	githubOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:8081/callback",
		Scopes:       []string{"user:email"},
	}
	oauthStateString = "random" // Replace with a secure random string
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type verifyRequest struct {
	Username string `json:"username"`
	Code     string `json:"code"`
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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

	if !isValidEmail(req.Username) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.Authenticate(req.Username, req.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "Error authenticating user", http.StatusInternalServerError)
		}
		return
	}

	if user == nil || !user.Verified {
		http.Error(w, "Invalid username or password, or email not verified", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := utils.GenerateJWT(req.Username)
	if err != nil {
		http.Error(w, "Error generating tokens", http.StatusInternalServerError)
		return
	}

	if err := h.UserService.UpdateRefreshToken(req.Username, refreshToken); err != nil {
		http.Error(w, "Error updating refresh token", http.StatusInternalServerError)
		return
	}

	res := tokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var req signupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if !isValidEmail(req.Username) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	if len(req.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	verificationCode := generateVerificationCode()

	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Verified: false,
		Code:     verificationCode,
	}

	if err := h.UserService.CreateUser(user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	if err := email.SendVerificationEmail(req.Username, verificationCode); err != nil {
		http.Error(w, "Error sending verification email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Verification email sent. Please check your inbox."))
}

func (h *AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	var req verifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.UserService.VerifyUser(req.Username, req.Code); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid verification codes", http.StatusBadRequest)
		} else {
			http.Error(w, "Error verifying user", http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte("Email verified successfully"))
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req refreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	claims, err := utils.ValidateJWT(req.RefreshToken)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	user, err := h.UserService.GetUserByUsername(claims.Username)
	if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if user == nil || user.RefreshToken != req.RefreshToken {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := utils.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generating tokens", http.StatusInternalServerError)
		return
	}

	if err := h.UserService.UpdateRefreshToken(user.Username, refreshToken); err != nil {
		http.Error(w, "Error updating refresh token", http.StatusInternalServerError)
		return
	}

	res := tokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (h *AuthHandler) HandleGitHubLogin(writer http.ResponseWriter, request *http.Request) {
	url := githubOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
}

func (h *AuthHandler) HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != oauthStateString {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := githubOauthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		log.Println("Code exchange failed:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken)
	if err != nil {
		log.Println("Failed to get user info:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer response.Body.Close()

	userInfo := make(map[string]interface{})
	json.NewDecoder(response.Body).Decode(&userInfo)

	log.Println("User info:", userInfo)

	// Here you can handle the user info, e.g., create a user session

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
