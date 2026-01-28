package validation

import (
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"user-management/internal/utils"

	"github.com/go-playground/validator/v10"
)

var (
	slugRegex           = regexp.MustCompile("^[a-z0-9]+(?:-[a-z0-9]+)*$")
	blockedDomainEmails = map[string]bool{
		"spam.com":     true,
		"edu.vn":       true,
		"tempmail.com": true,
	}
)

func RegisterCustomValidations(v *validator.Validate) {
	// Slug validation: lowercase letters, numbers, hyphens, no spaces, no special characters
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return slugRegex.MatchString(value)
	})

	// Advanced email validation: check for blocked domains
	v.RegisterValidation("email_advanced", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()

		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			return false
		}

		domain := utils.NormalizeString(parts[1])

		return !blockedDomainEmails[domain]
	})

	v.RegisterValidation("strong_password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 8 {
			return false
		}

		hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecial := regexp.MustCompile(`[!@#~$%^&*()+|_]{1,}`).MatchString(password)

		return hasLower && hasUpper && hasNumber && hasSpecial
	})

	// Minimum integer value validation
	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		min, err := strconv.ParseInt(fl.Param(), 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() >= min
	})

	// File extension validation
	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()

		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}
		allowedExts := strings.Split(allowedStr, " ")
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowedExts {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}

		return false
	})

	// Use JSON tags for field names in error messages
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
