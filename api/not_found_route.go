package api

import "github.com/gofiber/fiber/v2"

// NotFoundRoute - a function that sets a handler for non-existent routes
func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "sorry, endpoint is not found",
			})
		},
	)
}
