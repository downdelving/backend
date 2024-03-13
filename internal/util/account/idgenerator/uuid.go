package idgenerator

import (
	"github.com/google/uuid"
)

type UUID struct {
}

func (u *UUID) GenerateID() string {
	return uuid.New().String()
}
