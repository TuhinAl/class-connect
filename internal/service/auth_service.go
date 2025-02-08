package service

import (
	"golang-api/internal/repository"
	"golang-api/internal/utility/token"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(email, password string) (*token.LoginResponse, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type authService struct {
	studentRepo repository.StudentStore
}

type CustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
