package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomPassword(length int) (string, error) {
	password := make([]byte, length)
	printableChars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(printableChars))))
		if err != nil {
			return "", err
		}
		password[i] = printableChars[randomIndex.Int64()]
	}

	return string(password), nil
}
