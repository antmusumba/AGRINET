package pkg

import (
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultCost       = 12
	MinPasswordLength = 8
)

// Predefined error values
var (
	ErrEmailRequired     = errors.New("email is required")
	ErrPasswordRequired  = errors.New("password is required")
	ErrFirstNameRequired = errors.New("first name is required")
	ErrLastNameRequired  = errors.New("last name is required")
	ErrEmailExists       = errors.New("email already exists")
	ErrUserNotFound      = errors.New("user not found")
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	if err := validatePassword(password); err != nil {
		return "", err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", NewError(ErrorTypeInternal, "failed to hash password")
	}
	return string(bytes), nil
}

// CheckPassword checks if a plain password matches a hashed password
func CheckPassword(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return NewValidationError("password", "invalid password")
		}
		return NewError(ErrorTypeInternal, "failed to check password")
	}
	return nil
}

// validateEmail checks if the email format is valid
func validateEmail(email string) error {
	if email == "" {
		return ErrEmailRequired
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return NewValidationError("email", "invalid email format")
	}
	return nil
}

// validatePassword checks if the password meets security requirements
func validatePassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}
	if len(password) < MinPasswordLength {
		return NewValidationError("password", "password must be at least 8 characters")
	}
	return nil
}
