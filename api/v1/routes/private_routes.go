package routes

import (
	"github.com/c0caina/MyLibrary-fiber_rest_api/api/v1/handlers"
	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	// Post handler for /book route
	// This handler allows the user to create a new book in the database
	route.Post("/book", middleware.JWTProtected(), handlers.CreateBook) 
	// Put handler for /book route
	// This handler allows the user to update an existing book in the database
	route.Put("/book", middleware.JWTProtected(), handlers.UpdateBook)
	// Delete handler for /book route
	// This handler allows the user to delete an existing book from the database
	route.Delete("/book", middleware.JWTProtected(), handlers.DeleteBook)
}
