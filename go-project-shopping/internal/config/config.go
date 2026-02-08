package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	DatabaseUrl string
}

type Config struct {
	Port string
	Db   DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		Port: fmt.Sprintf(":%s", GetEnv("PORT", "8080")),
		Db: DatabaseConfig{
			DatabaseUrl: GetEnv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		},
	}
}

func GetEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
