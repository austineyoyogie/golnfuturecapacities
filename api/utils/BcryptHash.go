package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	const hashCost int = 12
	return bcrypt.GenerateFromPassword([]byte(password), hashCost)
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
