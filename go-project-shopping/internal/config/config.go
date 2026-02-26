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

type JwtConfig struct {
	SecretKey              string
	AccessTokenExpiration  string
	RefreshTokenExpiration string
}

type Config struct {
	Port          string
	EncryptionKey string
	AppEnv        string
	XApiKey       string

	Db    DatabaseConfig
	Redis RedisConfig
	Jwt   JwtConfig
}

var config Config

func NewConfig() {
	config = Config{
		Port:          fmt.Sprintf(":%s", utils.GetEnv("PORT", "8080")),
		AppEnv:        utils.GetEnv("APP_ENV", "development"),
		EncryptionKey: utils.GetEnv("ENCRYPTION_KEY", "ffdffafae19249232834372926bfefe7"),
		XApiKey:       utils.GetEnv("X_API_KEY", "21b8f79c-ba0e-485b-8a00-72b425a083a0"),

		Db: DatabaseConfig{
			DatabaseUrl: utils.GetEnv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		},
		Redis: RedisConfig{
			Address:  utils.GetEnv("REDIS_ADDRESS", "localhost:6379"),
			Username: utils.GetEnv("REDIS_USER", ""),
			Password: utils.GetEnv("REDIS_PASSWORD", ""),
			DB:       utils.GetEnvAsInt("REDIS_DB", 0),
		},
		Jwt: JwtConfig{
			SecretKey:              utils.GetEnv("JWT_SECRET_KEY", "your_secret_key"),
			AccessTokenExpiration:  utils.GetEnv("ACCESS_TOKEN_EXPIRATION", "15m"),
			RefreshTokenExpiration: utils.GetEnv("REFRESH_TOKEN_EXPIRATION", "168h"),
		},
	}
}

func Get() *Config {
	return &config
}
