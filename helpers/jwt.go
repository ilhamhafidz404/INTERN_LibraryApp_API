package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret_jwt_key")

func GenerateJWT(userID uint, nisn string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"nisn":    nisn,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expired 24 jam
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
