package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config stores all configuration of the application
type Config struct {
	// Server configuration
	ServerPort int
	ServerHost string
	
	// Database configuration
	DatabaseURL string
	
	// Logger configuration
	LogLevel string
	
	// JWT Secret
	JWTSecret string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()
	
	port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		port = 8080
	}
	
	return &Config{
		ServerPort:  port,
		ServerHost:  getEnv("SERVER_HOST", ""),
		DatabaseURL: getEnv("DATABASE_URL", "sqlite3://tasks.db"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
