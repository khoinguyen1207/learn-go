package utils

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
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
			case "required":
				errors[fieldErr.Field()] = fmt.Sprintf("%s is required", fieldErr.Field())
			default:
				errors[fieldErr.Field()] = fmt.Sprintf("%s is not valid", fieldErr.Field())
			}
		}

		return gin.H{"error": errors}
	}

	return gin.H{"error": err.Error()}
}
