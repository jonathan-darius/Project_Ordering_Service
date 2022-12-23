package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(hash)
}

func VerifyPassword(realPass string, proPass string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(proPass), []byte(realPass))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprint("Email or Password Incorrect")
		check = false
	}

	return check, msg
}
