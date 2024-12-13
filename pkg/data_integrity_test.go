package pkg

import (
	"log"
	"testing"
)

func TestPasswordHashingAndValidation(t *testing.T) {
	password := "securepassword123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Hashing failed: %v", err)
	}

	// Test with correct password
	err = CheckPassword(hashedPassword, password)
	if err != nil {
		t.Errorf("Password validation failed for correct password: %v", err)
	}

	// Test with incorrect password
	err = CheckPassword(hashedPassword, "wrongpassword")
	if err == nil {
		t.Error("Password validation succeeded for incorrect password")
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
