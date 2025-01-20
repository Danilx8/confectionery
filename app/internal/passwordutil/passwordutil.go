package passwordutil

import (
	"errors"
	"strings"
	"unicode"
)

func ValidateClientPassword(password string, login string) error {
	if len(password) < 5 || len(password) > 20 {
		return errors.New("password must be from 5 to 20 characters")
	}

	containsUpper := false
	containsLower := false

	for _, r := range password {
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			containsUpper = true
		} else if unicode.IsLetter(r) && unicode.IsLower(r) {
			containsLower = true
		}

		if containsLower && containsUpper {
			break
		}
	}

	if !(containsLower && containsUpper) {
		return errors.New("password must contain uppercase and lowercase letters")
	}

	if strings.Contains(strings.ToLower(password), strings.ToLower(login)) {
		return errors.New("password cannot contain login")
	}

	return nil
}
