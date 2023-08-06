package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

// JWTProtected is a function that returns a middleware handler for validating JWT tokens
// It uses jwtware to sign and verify tokens using the secret key from the environment variable JWT_SECRET_KEY
// It also uses jwtError to handle errors related to JWT tokens
func JWTProtected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ErrorHandler: jwtError,
	})
}

// jwtError is a function that handles errors related to JWT tokens
// It returns the appropriate status code and error message in JSON format
// It distinguishes two types of errors: missing or malformed token (400 Bad Request) and invalid token (401 Unauthorized)
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": err.Error(),
	})
}
