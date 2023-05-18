package util

import (
	"crypto/rand"
	"crypto/sha256"
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	return salt, err

}

func HashWithSalt(password string, salt []byte) []byte {
	salted := append(salt, []byte(password)...)
	hashed := sha256.Sum256(salted)
	return hashed[:]
}
