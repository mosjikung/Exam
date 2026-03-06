package config

import (
	"os"
)

// Config holds all application configuration.
type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Name string
	Port string
	Env  string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

// Load reads config from environment variables with sane defaults.
func Load() *Config {
	return &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "IT 06-1 Product API"),
			Port: getEnv("PORT", "3000"),
			Env:  getEnv("APP_ENV", "development"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "it06user"),
			Password: getEnv("DB_PASSWORD", "it06pass"),
			Name:     getEnv("DB_NAME", "it06db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			TimeZone: getEnv("DB_TIMEZONE", "Asia/Bangkok"),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
