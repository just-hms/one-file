package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {

	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hash), err
}

func Verify(hashedPassword string, plainPassword string) bool {

	bytePlain := []byte(plainPassword)
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)

	return err == nil
}
