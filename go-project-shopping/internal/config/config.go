package config

import (
	"fmt"
	"project-shopping/internal/utils"
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
		Port: fmt.Sprintf(":%s", utils.GetEnv("PORT", "8080")),
		Db: DatabaseConfig{
			DatabaseUrl: utils.GetEnv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		},
	}
}
