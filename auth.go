package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var googleClientID = os.Getenv("GOOGLE_CLIENT_ID")

func generateJWT(userID int, email string) (string, error) {
	var jwtSecretKey = os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}
