package routes

import (
	"github.com/c0caina/MyLibrary-fiber_rest_api/internal/app/controllers"
	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)

	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)

	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook)
}
