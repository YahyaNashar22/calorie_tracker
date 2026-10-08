package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	AppEnv string
	DBPath string
}

var AppConfig Config

func Load() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Failed to load .env file: %w", err)
	}

	AppConfig = Config{
		Port:   getEnv("APP_PORT", "4000"),
		AppEnv: getEnv("APP_ENV", "development"),
		DBPath: getEnv("DB_PATH", "./internal/database/calories.db"),
	}

	return nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
