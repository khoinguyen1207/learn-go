package validation

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("Failed to get validator engine")
	}

	RegisterCustomValidations(v)
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

	return gin.H{"error": fmt.Sprintf("Validation error: %s", err.Error())}
}
