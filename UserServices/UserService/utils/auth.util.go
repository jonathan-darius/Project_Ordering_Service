package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type SignedDetails struct {
	Email string
	Uid   string
	Role  string
	*jwt.RegisteredClaims
}

func GenerateToken(Email string, Uid string, uRole string) (signedToken string, signedRefreshToken string, err error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env File")
	}
	SecretKey := os.Getenv("SECRET_KEY")
	fmt.Println(Email, " ", Uid, " ", uRole)
	claims := SignedDetails{
		Email: Email,
		Uid:   Uid,
		Role:  uRole,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24))),
		},
	}
	refreshClaims := &SignedDetails{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(168)))},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SecretKey))
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env File")
	}
	SecretKey := os.Getenv("SECRET_KEY")
	tokenString := signedToken
	token, err := jwt.ParseWithClaims(
		tokenString,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		return
	}
	return claims, msg
}
