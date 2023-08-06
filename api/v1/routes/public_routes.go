package routes

import (
	"github.com/c0caina/MyLibrary-fiber_rest_api/database/postgres"
	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PublicRoutes(a *fiber.App) {

	route := a.Group("/api/v1")
	// Get handler for /books route
	// This handler allows the user to get all the books from the database
	route.Get("/books", func (c *fiber.Ctx) error {
		// Connect to PostgreSQL database
		db, err := postgres.NewPostgres()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Get all the books from the database
		books, err := db.GetBooks()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Return JSON response with the books
        // We also include the count of the books for convenience
		return c.JSON(fiber.Map{
			"error": err,
			"count": len(books),
			"books": books,
		})
	})
	// Get handler for /book/:id route
	// This handler allows the user to get a specific book from the database by id
	route.Get("/book/:id", func (c *fiber.Ctx) error {
		// Parse the id parameter from the request as a UUID 
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Connect to PostgreSQL database
		db, err := postgres.NewPostgres()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Get the book from the database by id 
		book, err := db.GetBook(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
				"book":  book,
			})
		}
		// Return JSON response with the book
		return c.JSON(fiber.Map{
			"error": err,
			"book":  book,
		})
	})
	// Get handler for /token/new route
	// This handler allows the user to generate a new access token using JWT
	route.Get("/token/new", func (c *fiber.Ctx) error {
		// Generate a new access token using the utils package
		token, err := utils.GenerateAccessJWT()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Return JSON response with the access token
		return c.JSON(fiber.Map{
			"error":        err,
			"access_token": token,
		})
	})
}
