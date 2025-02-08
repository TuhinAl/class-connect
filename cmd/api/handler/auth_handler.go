package handler

import (
	"encoding/json"
	"golang-api/internal/service"
	"golang-api/internal/utility/token"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq token.LoginRequest

	// Decode the JSON request body into the login request struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the authService to attempt login with the provided credentials
	response, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Set content type header and encode response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error processing response", http.StatusInternalServerError)
	}
}
