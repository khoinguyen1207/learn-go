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

var sensitiveFields = map[string]bool{
	"password":              true,
	"password_confirmation": true,
	"confirm_password":      true,
	"old_password":          true,
	"new_password":          true,
	"secret":                true,
	"secret_key":            true,
	"api_key":               true,
	"access_token":          true,
	"refresh_token":         true,
	"token":                 true,
	"authorization":         true,
	"credit_card":           true,
	"card_number":           true,
	"cvv":                   true,
	"ssn":                   true,
	"private_key":           true,
}

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

func GetEnvAsSlice(name string, defaultValue []string, sep string) []string {
	if valueStr := os.Getenv(name); valueStr != "" {
		return strings.Split(valueStr, sep)
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

func SanitizeBody(data map[string]any) map[string]any {
	sanitized := make(map[string]any, len(data))

	for key, value := range data {
		lowerKey := strings.ToLower(key)

		if sensitiveFields[lowerKey] {
			sanitized[key] = "***REDACTED***"
			continue
		}

		switch v := value.(type) {
		case map[string]any:
			// Nested map, sanitize recursively
			sanitized[key] = SanitizeBody(v)
		case []any:
			// Handle slice
			sanitized[key] = sanitizeSlice(v)
		default:
			sanitized[key] = value
		}
	}

	return sanitized
}

func sanitizeSlice(slice []any) []any {
	result := make([]any, len(slice))
	for i, item := range slice {
		switch v := item.(type) {
		case map[string]any:
			result[i] = SanitizeBody(v)
		case []any:
			result[i] = sanitizeSlice(v)
		default:
			result[i] = item
		}
	}
	return result
}

func GenerateRandomString(byteLength int) (string, error) {
	nonce := make([]byte, byteLength)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(nonce), nil
}
