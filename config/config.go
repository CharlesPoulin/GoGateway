package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string
	DBConnectionStr string
	ExternalAPIBase string
	RequestTimeout  time.Duration
	LogLevel        string
}

// LoadConfig loads configuration from .env file and environment variables
func LoadConfig() *Config {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Construct the PostgreSQL connection string dynamically
	dbUser := getEnv("DB_USER", "")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")
	dbConnectionStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=" + dbSSLMode

	config := &Config{
		ServerPort:      getEnv("SERVER_PORT", ":3000"),
		DBConnectionStr: dbConnectionStr,
		ExternalAPIBase: getEnv("EXTERNAL_API_BASE", "http://localhost:4001"),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
	}

	// Parse duration for request timeout
	timeoutStr := getEnv("REQUEST_TIMEOUT", "60s")
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		log.Fatalf("Invalid REQUEST_TIMEOUT: %v", err)
	}
	config.RequestTimeout = timeout

	return config
}

// Helper function to get an environment variable or default value
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
