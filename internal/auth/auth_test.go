package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {
	id := uuid.New()
	secret := "my-ultra-secret-key"

	// Тест создания и валидации
	token, err := MakeJWT(id, secret, time.Hour)
	if err != nil {
		t.Fatalf("Error making JWT: %v", err)
	}

	parsedID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("Error validating JWT: %v", err)
	}

	if parsedID != id {
		t.Errorf("Expected ID %v, got %v", id, parsedID)
	}
}
