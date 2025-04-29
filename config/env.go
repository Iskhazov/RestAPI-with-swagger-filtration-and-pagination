package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	PublicHost            string
	DBUser                string
	DBPass                string
	Port                  string
	DBAddr                string
	DBName                string
	JWTExprirationSeconds int64
	JWTSecret             string
}

var Envs = InitConfig()

func InitConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: GetEnv("PUBLIC_HOST", "localhost"),
		DBUser:     GetEnv("DB_USER", "postgres"),
		DBPass:     GetEnv("DB_PASSWORD", "loloshka777"),
		Port:       GetEnv("PORT", "5432"),
		DBName:     GetEnv("DB_NAME", "user_initials"),
		//JWTSecret:             GetEnv("JWT_SECRET", ""),
		//JWTExprirationSeconds: GetEnvAsInt("JWT_EXP", 3600*24*7),
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback
	}
}

func GetEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	} else {
		return fallback
	}
}
