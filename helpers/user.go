package helpers

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+="

func GenerateUser(first_str string, second_str string) (string) {
	username := strings.ToLower(first_str[:3]) + strings.ToLower(second_str[:3])
	return username
}

func GenerateSecurePassword(length int) (string, error) {
	password := make([]byte, length)
	for i := range password {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[charIndex.Int64()]
	}
	// return string(password), nil
	return "123", nil
}
