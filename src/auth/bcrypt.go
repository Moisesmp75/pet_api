package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt_password(password string) string {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(err.Error())
	}
	return string(encrypt)
}

func DecryptPasswordHash(hash string, password string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return result == nil
}