package routes

import (
	"time"

	"github.com/c0caina/MyLibrary-fiber_rest_api/database/postgres"
	"github.com/c0caina/MyLibrary-fiber_rest_api/entities"
	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/middleware"
	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	// Post handler for /book route
	// This handler allows the user to create a new book in the database
	route.Post("/book", middleware.JWTProtected(), func (c *fiber.Ctx) error {
		// Parse JWT from request
		claims, err := utils.ParceJWT(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Check if token is expired and return unauthorized status if so
		if time.Now().Unix() > claims.Expires {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized, check expiration time of your token",
			})
		}
		// Create a new book struct and parse request body into it
		book := &entities.Book{}
		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
		// Validate the book struct using a validator
		// This is necessary because the data comes from the user and may be invalid or malicious
		if err := utils.NewValidator().Struct(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ValidatorErrors(err),
			})
		}
		// Create the book in the database
		if err := db.CreateBook(book); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Return JSON response with the book
		return c.JSON(fiber.Map{
			"error": err,
			"book":  book,
		})
	})
	// Put handler for /book route
	// This handler allows the user to update an existing book in the database
	route.Put("/book", middleware.JWTProtected(), func (c *fiber.Ctx) error {
		// Parse JWT from request
		claims, err := utils.ParceJWT(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Check if token is expired and return unauthorized status if so
		if time.Now().Unix() > claims.Expires {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized, check expiration time of your token",
			})
		}
		// Create a new book struct and parse request body into it
		book := &entities.Book{}
		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
		// Validate the book struct using a validator 
        // This is necessary because the data comes from the user and may be invalid or malicious
		if err := utils.NewValidator().Struct(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ValidatorErrors(err),
			})
		}
		// Update the book in the database 
		if err := db.UpdateBook(book); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Return status code 201 (created) to indicate success
		return c.SendStatus(fiber.StatusCreated)
	})
	// Delete handler for /book route
	// This handler allows the user to delete an existing book from the database
	route.Delete("/book", middleware.JWTProtected(), func (c *fiber.Ctx) error {
		// Parse JWT from request
		claims, err := utils.ParceJWT(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Check if token is expired and return unauthorized status if so
		if time.Now().Unix() > claims.Expires {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized, check expiration time of your token",
			})
		}
		// Create a new book struct and parse request body into it
		book := &entities.Book{}
		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Validate the book struct using a validator 
        // This is necessary because the data comes from the user and may be invalid or malicious
        // We only validate the id field because it is the only one we need to delete the book
		if err := utils.NewValidator().StructPartial(book, "id"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ValidatorErrors(err),
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
        // If no book is found, return not found status
		foundedBook, err := db.GetBook(book.ID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "book with this ID not found",
			})
		}
		// Delete the book from the database by id 
		if err := db.DeleteBook(foundedBook.ID); err != nil {
	
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// Return status code 204 (no content) to indicate success
		return c.SendStatus(fiber.StatusNoContent)
	})
}
