package utils

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func NormalizeString(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
