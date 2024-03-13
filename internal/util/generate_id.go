package util

import (
	"github.com/google/uuid"
)

// Generate a UUID.
func GenerateID() string {
	return uuid.New().String()
}
