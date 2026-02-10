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

var RATE_LIMIT_REQUEST_SECOND int
var RATE_LIMIT_REQUEST_BURST int

func NewConfig() *Config {
	// Load rate limit settings from environment variables
	RATE_LIMIT_REQUEST_SECOND = utils.GetEnvAsInt("RATE_LIMIT_REQUEST_SECOND", 5)
	RATE_LIMIT_REQUEST_BURST = utils.GetEnvAsInt("RATE_LIMIT_REQUEST_BURST", 10)

	return &Config{
		Port: fmt.Sprintf(":%s", utils.GetEnv("PORT", "8080")),
		Db: DatabaseConfig{
			DatabaseUrl: utils.GetEnv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		},
	}
}
