package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	Environment string
}

func LoadConfig() *Config {
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
