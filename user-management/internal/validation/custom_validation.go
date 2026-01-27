package validation

import (
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations(v *validator.Validate) {
	// Slug validation: lowercase letters, numbers, hyphens, no spaces, no special characters
	slugRegex := regexp.MustCompile("^[a-z0-9]+(?:-[a-z0-9]+)*$")
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return slugRegex.MatchString(value)
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
