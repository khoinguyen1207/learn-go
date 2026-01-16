package utils

import (
	"fmt"
	"regexp"

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
