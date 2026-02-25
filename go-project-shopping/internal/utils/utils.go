package utils

import (
	"os"
	"project-shopping/pkg/logger"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

func NormalizeString(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func GetEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func GetEnvAsInt(name string, defaultValue int) int {
	if valueStr := os.Getenv(name); valueStr != "" {
		intVal, err := strconv.Atoi(valueStr)
		if err != nil {
			return defaultValue
		}
		return intVal
	}
	return defaultValue
}

func NewLoggerWithPath(path string, level string) *zerolog.Logger {
	config := logger.LoggerConfig{
		Level:      level,
		Filename:   path,
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     7, // days
		Compress:   true,
		IsDev:      GetEnv("APP_ENV", "production"),
	}

	return logger.NewLoggerConfig(config)
}

func SanitizeBody(data map[string]any) {
	sensitive := []string{"password", "token", "secret"}
	for _, key := range sensitive {
		if _, ok := data[key]; ok {
			data[key] = "****"
		}
	}
}
