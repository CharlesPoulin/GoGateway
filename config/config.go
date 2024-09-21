package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	ServerPort      string
	DBConnectionStr string
	ExternalAPIBase string
	RequestTimeout  time.Duration
	LogLevel        string
	// Add other configuration fields as needed
}

func LoadConfig() *Config {
	config := &Config{
		ServerPort:      getEnv("SERVER_PORT", ":3000"),
		DBConnectionStr: getEnv("DB_CONNECTION_STR", "user=youruser password=yourpassword dbname=yourdb sslmode=disable"),
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

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
