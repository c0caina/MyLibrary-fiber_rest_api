package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAccessJWT is a function that generates a JWT token for accessing protected resources
// It uses jwt to create and sign the token using the secret key from the environment variable JWT_SECRET_KEY
// It also sets the expiration time of the token according to the environment variable JWT_SECRET_KEY_EXPIRE_HOURS_COUNT
// It returns the token as a string or an error if something went wrong
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
