//go:build testing

package passwordhasher

type Identity struct {
}

func (i *Identity) HashPassword(password string) (string, error) {
	return password, nil
}

func (i *Identity) CheckPassword(hashedPassword, password string) bool {
	return hashedPassword == password
}
