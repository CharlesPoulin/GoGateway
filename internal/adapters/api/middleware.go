package api

import (
	"context"
	"net/http"
	"strings"

	"GoGateway/internal/app"
	"GoGateway/util"
)

type Middleware struct {
	AuthService app.AuthService
	Logger      util.Logger
}

// Context key for storing user data (you can define this globally)
type contextKey string

const (
	ContextUserKey = contextKey("user")
)

// NewMiddleware creates a new middleware instance
func NewMiddleware(authService app.AuthService, logger util.Logger) *Middleware {
	return &Middleware{
		AuthService: authService,
		Logger:      logger,
	}
}

// AuthMiddleware validates JWT tokens and adds user information to the request context
func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			util.RespondWithError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		// Split the "Authorization" header into "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			util.RespondWithError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		tokenString := parts[1]

		// Validate the token and extract claims
		claims, valid, err := m.AuthService.ValidateAndExtractClaims(tokenString)
		if err != nil || !valid {
			util.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Optionally log the authenticated user's username or any other claim
		m.Logger.Info("Authenticated user: " + claims["username"].(string))

		// Pass the user information to the next handler via context
		ctx := context.WithValue(r.Context(), ContextUserKey, claims["username"].(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
