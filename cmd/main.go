package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"GoGateway/config"
	"GoGateway/infra"
	"GoGateway/infra/db"
	"GoGateway/internal/adapters/api"
	"GoGateway/internal/app"
	"GoGateway/util"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Logger
	logger := util.NewLogger(cfg.LogLevel)
	logger.Info("Starting API Gateway")

	// Initialize Database Repository
	authRepo, err := db.NewAuthRepository(cfg.DBConnectionStr, logger)
	if err != nil {
		logger.Fatal("Failed to initialize Auth Repository", "error", err)
	}
	defer authRepo.Close()

	// Initialize HTTP Client for External APIs
	httpClient := infra.NewHTTPClient(cfg.RequestTimeout, logger)

	// Initialize Application Service
	authService := app.NewAuthService(authRepo, httpClient, cfg.ExternalAPIBase, logger)

	// Initialize API Handlers
	handlers := api.NewHandler(authService, logger)

	// Initialize Router
	router := api.NewRouter(handlers, logger)

	// Initialize HTTP Server
	server := &http.Server{
		Addr:    cfg.ServerPort,
		Handler: router,
		// Timeouts to prevent Slowloris attacks
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Start Server in a Goroutine
	go func() {
		logger.Info("API Gateway is listening", "port", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("ListenAndServe failed", "error", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", "error", err)
	}

	logger.Info("Server exiting")
}
