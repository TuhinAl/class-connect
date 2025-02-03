package security

import (
	"golang-api/models"

	"github.com/golang-jwt/jwt"
)

var (
    JWTSecretKey = []byte("your-secret-key") // In production, use environment variable
)

type AuthService interface {
    Login(email, password string) (*models.LoginResponse, error)
    ValidateToken(tokenString string) (*jwt.Token, error)
}

