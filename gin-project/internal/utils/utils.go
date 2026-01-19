package utils

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func IsValidID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func RegexValidate(fieldName string, value string, re *regexp.Regexp) error {
	if !re.MatchString(value) {
		return fmt.Errorf("Invalid %s format", fieldName)
	}

	return nil
}

func HandleValidationError(err error) gin.H {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, fieldErr := range validationErrors {
			switch fieldErr.Tag() {
			case "gt":
				errors[fieldErr.Field()] = fmt.Sprintf("%s must be greater than %s", fieldErr.Field(), fieldErr.Param())
			case "lt":
				errors[fieldErr.Field()] = fmt.Sprintf("%s must be less than %s", fieldErr.Field(), fieldErr.Param())
			case "required":
				errors[fieldErr.Field()] = fmt.Sprintf("%s is required", fieldErr.Field())
			case "slug":
				errors[fieldErr.Field()] = fmt.Sprintf("%s is not a valid", fieldErr.Field())
			case "min":
				errors[fieldErr.Field()] = fmt.Sprintf("%s must be at least %s characters long", fieldErr.Field(), fieldErr.Param())
			case "max":
				errors[fieldErr.Field()] = fmt.Sprintf("%s must be at most %s characters long", fieldErr.Field(), fieldErr.Param())
			case "oneof":
				allowed := strings.Join(strings.Split(fieldErr.Param(), " "), ", ")
				errors[fieldErr.Field()] = fmt.Sprintf("%s must be one of the following: %s", fieldErr.Field(), allowed)
			case "min_int":
				errors[fieldErr.Field()] = fmt.Sprintf("%s must be at least %s", fieldErr.Field(), fieldErr.Param())
			case "file_ext":
				allowed := strings.Join(strings.Split(fieldErr.Param(), " "), ",")
				errors[fieldErr.Field()] = fmt.Sprintf("%s must have one of the following extensions: %s", fieldErr.Field(), allowed)
			default:
				errors[fieldErr.Field()] = fmt.Sprintf("%s is not valid", fieldErr.Field())
			}
		}

		return gin.H{"error": errors}
	}

	return gin.H{"error": err.Error()}
}

func RegisterCustomValidations() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	slugRegex := regexp.MustCompile("^[a-z0-9]+(?:-[a-z0-9]+)*$")
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return slugRegex.MatchString(value)
	})

	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		min, err := strconv.ParseInt(fl.Param(), 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() >= min
	})

	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()

		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}
		allowedExts := strings.Split(allowedStr, " ")
		// filename = .png
		// toLower = .png
		// TrimPrefix = png
		// allowedExts = [jpg, png, jpeg]
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowedExts {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}

		return false
	})

	return nil
}
