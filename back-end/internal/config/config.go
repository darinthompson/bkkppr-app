package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds all configuration variables
type Config struct {
	Port        string
	DatabaseURL string
}

// LoadConfig reads .env file and loads configuration
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using defaults")
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "root:root@tcp(localhost:3306)/bkkppr?charset=utf8mb4&parseTime=True&loc=Local"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
