package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	validateUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	validateFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain length of %d-%d", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(username string) error {

	err := ValidateString(username, 3, 20)
	if err != nil {
		return err
	}

	if !validateUsername(username) {
		return fmt.Errorf("must contain only letters, digits or underscore ")
	}

	return nil
}

func ValidateFullname(fullname string) error {

	err := ValidateString(fullname, 3, 100)
	if err != nil {
		return err
	}

	if !validateFullName(fullname) {
		return fmt.Errorf("must contain only letters and white spaces ")
	}

	return nil
}

func ValidatePassword(password string) error {
	return ValidateString(password, 6, 100)
}

func ValidateEmail(email string) error {
	err := ValidateString(email, 6, 100)
	if err != nil {
		return err
	}

	_, err = mail.ParseAddress(email)

	if err != nil {
		return fmt.Errorf("invalid email")
	}

	return nil
}
