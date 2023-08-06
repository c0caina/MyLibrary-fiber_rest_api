package main

import (
	"log"
	"os"

	"github.com/c0caina/MyLibrary-fiber_rest_api/api"
	"github.com/c0caina/MyLibrary-fiber_rest_api/api/v1/routes"
	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

//Todo: Swagger частично сломан в PUT, POST, DELETE. Не может корректно прочитать запрос.

// @title MyLibrary API
// @version 0.1
// @description Documentation with the ability to touch the api.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	app := fiber.New()

	middleware.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	api.NotFoundRoute(app)
	if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Server is not running! Error: %v", err)
	}
}
