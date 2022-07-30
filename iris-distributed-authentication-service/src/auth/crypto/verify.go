package crypto

import "golang.org/x/crypto/bcrypt"

func Verify(encode, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encode), []byte(password))

	return err == nil
}
