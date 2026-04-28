package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("super_secret_key")

func GenerateToken(userID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
