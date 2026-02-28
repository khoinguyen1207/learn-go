package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"path/filepath"
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

func NewLoggerWithPath(filename string, level string) *zerolog.Logger {
	dir := GetRootDir()

	filepath := filepath.Join(dir, "internal/logs", filename)
	config := logger.LoggerConfig{
		Level:      level,
		Filename:   filepath,
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     7, // days
		Compress:   true,
		IsDev:      GetEnv("APP_ENV", "production"),
	}

	return logger.NewLogger(config)
}

func GetRootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Error getting root directory")
	}
	return dir
}

func SanitizeBody(data map[string]any) {
	sensitive := []string{"password", "token", "secret"}
	for _, key := range sensitive {
		if _, ok := data[key]; ok {
			data[key] = "****"
		}
	}
}

func GenerateRandomString(byteLength int) (string, error) {
	nonce := make([]byte, byteLength)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(nonce), nil
}
