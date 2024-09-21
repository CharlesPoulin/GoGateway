package api

import (
	"net/http"
	"strings"

	"GoGateway/internal/app"
	"GoGateway/util"
)

type Middleware struct {
	AuthService app.AuthService
	Logger      util.Logger
}

func NewMiddleware(authService app.AuthService, logger util.Logger) *Middleware {
	return &Middleware{
		AuthService: authService,
		Logger:      logger,
	}
}

// AuthMiddleware validates JWT tokens
func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			util.RespondWithError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			util.RespondWithError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		token := parts[1]
		valid, err := m.AuthService.ValidateToken(token)
		if err != nil || !valid {
			util.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
