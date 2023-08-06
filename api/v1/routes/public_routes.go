package routes

import (
	"github.com/c0caina/MyLibrary-fiber_rest_api/api/v1/handlers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {

	route := a.Group("/api/v1")
	
	// Get handler for /books route
	// This handler allows the user to get all the books from the database
	// GetBooks func gets all exists books.
	route.Get("/books", handlers.GetBooks)
	// Get handler for /book/:id route
	// This handler allows the user to get a specific book from the database by id
	route.Get("/book/:id", handlers.GetBook)
	// Get handler for /token/new route
	// This handler allows the user to generate a new access token using JWT
	route.Get("/token/new", handlers.GetNewJWT)
}
