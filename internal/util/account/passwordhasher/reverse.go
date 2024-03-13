//go:build testing

package passwordhasher

type Reverse struct {
}

func (r *Reverse) HashPassword(password string) (string, error) {
	runes := []rune(password)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}

func (r *Reverse) ComparePassword(hashedPassword, password string) (bool, error) {
	return hashedPassword == password, nil
}

func (r *Reverse) CheckPassword(hashedPassword, password string) bool {
	return hashedPassword == password
}
