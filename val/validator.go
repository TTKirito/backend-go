package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
)

func validateString(value string, minLength int, maxLength int) error {
	n := len(value)

	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}

	return nil
}

func ValidateUsername(value string) error {
	if err := validateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, digits, underscore")
	}

	return nil

}

func ValidateFullName(value string) error {
	if err := validateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidFullName(value) {
		return fmt.Errorf("must contain only leters or space")
	}

	return nil
}

func ValidatePassword(value string) error {
	return validateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := validateString(value, 3, 200); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}

	return nil
}
