package pkg

import (
	"log"
	"testing"
)

func TestPasswordHashingAndValidation(t *testing.T) {
	plainPassword := "securepassword123"
	wrongPassword := "wrongpassword"

	// Hash the password
	hashedPassword, err := HashPassword(plainPassword)
	if err != nil {
		t.Fatalf("Hashing failed: %v", err)
	}

	// Log the hashed password for debugging (optional)
	log.Printf("Generated Hash: %s", hashedPassword)

	// Validate the correct password
	if err := CheckPassword(hashedPassword, plainPassword); err != nil {
		t.Errorf("Validation failed for correct password: %v", err)
	}

	// Validate an incorrect password
	if err := CheckPassword(hashedPassword, wrongPassword); err == nil {
		t.Error("Validation succeeded for incorrect password")
	}
}

func TestPasswordHashingValidation(t *testing.T) {
	plainPassword := "mypassword123"

	hashedPassword, err := HashPassword(plainPassword)
	if err != nil {
		t.Fatalf("Hashing failed: %v", err)
	}
	log.Printf("Generated Hash: %s", hashedPassword)

	// Test with the correct password
	if err := CheckPassword(hashedPassword, plainPassword); err != nil {
		t.Errorf("Password validation failed: %v", err)
	}

	// Test with an incorrect password
	if err := CheckPassword(hashedPassword, "wrongpassword"); err == nil {
		t.Error("Validation succeeded for incorrect password")
	}
}
