package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessJWT() (string, error) {

	hoursCount, err := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_HOURS_COUNT"))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
