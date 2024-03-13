package util

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) (bool, error)
	CheckPassword(hashedPassword, password string) bool
}
