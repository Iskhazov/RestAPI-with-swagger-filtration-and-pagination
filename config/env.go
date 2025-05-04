package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBAddr string
	DBName string
	DBPort string
}

var Envs = InitConfig()

func InitConfig() Config {
	godotenv.Load()

	return Config{

		DBUser: GetEnv("DB_USER", "postgres"),
		DBPass: GetEnv("DB_PASSWORD", "loloshka777"),
		DBHost: GetEnv("DB_HOST", "localhost"),
		DBName: GetEnv("DB_NAME", "user_initials"),
		DBPort: GetEnv("DB_PORT", "5432"),
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback
	}
}
