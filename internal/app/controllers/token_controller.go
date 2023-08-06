package controllers

import (
	utils "github.com/c0caina/MyLibrary-fiber_rest_api/pkg/utilities"
	"github.com/gofiber/fiber/v2"
)

func GetNewAccessToken(c *fiber.Ctx) error {
	token, err := utils.GenerateAccessJWT()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
