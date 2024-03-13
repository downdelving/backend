package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Cost of hashing the password.
const cost = 0

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
