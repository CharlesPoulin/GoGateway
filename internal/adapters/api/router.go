package api

import (
	"net/http"
	"time"

	"GoGateway/util"

	"github.com/go-chi/chi/middleware" // Add the correct middleware import
	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *Handler, logger util.Logger) http.Handler {
	r := chi.NewRouter()

	// Global Middlewares
	r.Use(middleware.RequestID)                 // Use the correct middleware
	r.Use(middleware.RealIP)                    // Use the correct middleware
	r.Use(middleware.Logger)                    // Use the correct middleware
	r.Use(middleware.Recoverer)                 // Use the correct middleware
	r.Use(middleware.Timeout(60 * time.Second)) // Use the correct middleware

	// Initialize API-specific Middleware
	apiMiddleware := NewMiddleware(handler.AuthService, logger)

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Post("/authenticate", handler.Authenticate)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			util.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Welcome to the API Gateway"})
		})
	})

	// Protected Routes
	r.Group(func(r chi.Router) {
		r.Use(apiMiddleware.Auth)

		r.Get("/user", handler.GetUser)
		// Add other protected routes here
	})

	// Health Check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		util.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "OK"})
	})

	return r
}
