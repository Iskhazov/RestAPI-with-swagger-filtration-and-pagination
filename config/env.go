package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

// Config holds the database configuration loaded from environment variables
type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBAddr string
	DBName string
	DBPort string
}

// Envs is the exported initialized configuration
var Envs = InitConfig()

// InitConfig loads environment variables from .env file and returns Config
func InitConfig() Config {
	// Specify path to .env file located in the 'envs' directory
	envFile := filepath.Join("envs", ".env")

	err := godotenv.Load(envFile)
	if err != nil {
		log.Println("Warning: .env file not found in 'envs' folder, relying on system environment variables")
	}

	// Return populated config struct
	return Config{
		DBUser: GetEnv("DB_USER"),
		DBPass: GetEnv("DB_PASSWORD"),
		DBHost: GetEnv("DB_HOST"),
		DBName: GetEnv("DB_NAME"),
		DBPort: GetEnv("DB_PORT"),
	}
}

// GetEnv fetches the value of an environment variable or exits if not found
func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}
