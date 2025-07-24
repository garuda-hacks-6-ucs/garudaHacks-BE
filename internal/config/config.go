package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds all configuration for the application.
type Config struct {
	ServerPort  string
	MongoURI    string
	MongoDbName string
}

// Load loads config from .env file
func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDbName: getEnv("MONGO_DB_NAME", "govtech"),
	}, nil
}

// Helper function to get env variable or return default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
