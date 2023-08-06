package routes

import (
	_ "github.com/c0caina/MyLibrary-fiber_rest_api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoute(a *fiber.App) {
	a.Get("/swagger/api/v1/*", swagger.HandlerDefault)
}