package utils

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

func NewValidator() *validator.Validate {
	validatorRequest := validator.New()

	if err := validatorRequest.RegisterValidation("password", PasswordValidator); err != nil {
		return nil
	}

	if err := validatorRequest.RegisterValidation("phone", PhoneValidator); err != nil {
		return nil
	}

	return validatorRequest
}

func PhoneValidator(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	if phoneNumber == "" {
		return true
	}

	parsedNumber, err := phonenumbers.Parse(phoneNumber, "VN")
	if err != nil {
		return false
	}

	return phonenumbers.IsValidNumber(parsedNumber)
}

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	// Check for minimum length (e.g., 8 characters)
	if len(password) < 8 {
		return false
	}

	// Check for at least one uppercase letter
	if !containsUppercase(password) {
		return false
	}

	// Check for at least one lowercase letter
	if !containsLowercase(password) {
		return false
	}

	// Check for at least one digit
	if !containsDigit(password) {
		return false
	}

	// Check for at least one special character (e.g., !@#$%^&*)
	if !containsSpecialCharacter(password) {
		return false
	}

	return true
}

// Helper functions to check for character types
func containsUppercase(s string) bool {
	return strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func containsLowercase(s string) bool {
	return strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz")
}

func containsDigit(s string) bool {
	return strings.ContainsAny(s, "0123456789")
}

func containsSpecialCharacter(s string) bool {
	// Define a regular expression to match special characters
	specialCharacterPattern := `[!@#$%^&*()_+{}\[\]:;<>,.?~]`
	match, _ := regexp.MatchString(specialCharacterPattern, s)
	return match
}
