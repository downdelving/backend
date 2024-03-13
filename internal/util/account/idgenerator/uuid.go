package idgenerator

import (
	"github.com/google/uuid"
)

type Uuid struct {
}

// Generate a UUID.
func (u *Uuid) GenerateId() string {
	return uuid.New().String()
}
