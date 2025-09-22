package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret string
}

type ServerConfig struct {
	Port    string
	GinMode string
}

func Load() *Config {
	if err := godotenv.Load("config.env"); err != nil {
		log.Println("Warning: config.env file not found, using system environment variables")
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "123456"),
			Name:     getEnv("DB_NAME", "blog"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "ktjnCkMI6GgMN3w6Nein+BFSl7YThGzlmwuomDSvkzo="),
		},
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
