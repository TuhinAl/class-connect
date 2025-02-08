package middleware

import (
	"context"
	"golang-api/internal/security"
	"net/http"
	"strings"
)

func AuthMiddleware(authService security.AuthService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "Authorization header is required", http.StatusUnauthorized)
                return
            }

            // Extract token from "Bearer <token>"
            tokenParts := strings.Split(authHeader, " ")
            if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
                http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
                return
            }

            token, err := authService.ValidateToken(tokenParts[1])
            if err != nil || !token.Valid {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            claims, ok := token.Claims.(*security.CustomClaims)
            if !ok {
                http.Error(w, "Invalid token claims", http.StatusUnauthorized)
                return
            }

            // Add user info to the context
            ctx := context.WithValue(r.Context(), "user_id", claims.ID)
            ctx = context.WithValue(ctx, "email", claims.Email)
            ctx = context.WithValue(ctx, "role", claims.Role)

            // Pass context into the next handler
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

type CustomClaims struct {
    UserID string
    Email  string
    Role   string
    // Add other claim fields if needed
}