package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassCode(passCode string) (string, error) {
	hashedPassCode, err := bcrypt.GenerateFromPassword([]byte(passCode), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassCode), nil
}

func CheckPassCode(passCode, hashedPassCode string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassCode), []byte(passCode))
	return err == nil
}
