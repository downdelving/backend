package passwordhasher

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	cost int
}

const cost = 14

func NewBcrypt() *Bcrypt {
	return &Bcrypt{cost: cost}
}

func (b *Bcrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	return string(bytes), err
}

func (b *Bcrypt) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
