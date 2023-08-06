package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FiberMiddleware is a function that adds some middleware handlers to the Fiber app
// It uses logger to log HTTP requests and responses
// This helps to improve the security and debugging of the app
func FiberMiddleware(a *fiber.App) {
	a.Use(
		logger.New(),
	)
}
