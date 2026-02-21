package config

import (
	"fmt"
	"project-shopping/internal/utils"
)

type DatabaseConfig struct {
	DatabaseUrl string
}

type RedisConfig struct {
	Address  string
	Username string
	Password string
	DB       int
}

type Config struct {
	Port  string
	Db    DatabaseConfig
	Redis RedisConfig
}

func NewConfig() *Config {
	return &Config{
		Port: fmt.Sprintf(":%s", utils.GetEnv("PORT", "8080")),
		Db: DatabaseConfig{
			DatabaseUrl: utils.GetEnv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		},
		Redis: RedisConfig{
			Address:  utils.GetEnv("REDIS_ADDRESS", "localhost:6379"),
			Username: utils.GetEnv("REDIS_USER", ""),
			Password: utils.GetEnv("REDIS_PASSWORD", ""),
			DB:       utils.GetEnvAsInt("REDIS_DB", 0),
		},
	}
}
