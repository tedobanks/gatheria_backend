package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func CheckPasswordHash(hashPassword, stringPassword  string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(stringPassword))
	return err == nil
}
