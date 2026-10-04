package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	Environment string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("[Config Info] No .env file found or error loading .env, falling back to OS environment variables.")
	} else {
		log.Println("[Config] Successfully loaded environment variables from .env file.")
	}

	port := getEnv("PORT", "8080")
	env := getEnv("APP_ENV", "development")

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "buma_db")
	dbSSL := getEnv("DB_SSLMODE", "disable")

	defaultConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSSL)

	dbURL := getEnv("DATABASE_URL", defaultConnStr)

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
		Environment: env,
	}
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
