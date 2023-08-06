package main

import (
	"log"
	"os"

	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/middleware"
	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}
