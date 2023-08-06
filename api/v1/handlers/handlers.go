package handlers

import (
	"time"

	"github.com/c0caina/MyLibrary-fiber_rest_api/database/postgres"
	"github.com/c0caina/MyLibrary-fiber_rest_api/entities"
	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetBooks func gets all exists books.
// @Summary Get all books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} entities.Book
// @Router /books [get]
func GetBooks(c *fiber.Ctx) error {
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
}

// GetBook func gets book by given ID or 404 error.
// @Summary Get book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} entities.Book
// @Router /book/{id} [get]
func GetBook(c *fiber.Ctx) error {
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
}

// GetNewJWT method for create a new access token.
// @Summary Create a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /token/new [get]
func GetNewJWT(c *fiber.Ctx) error {
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
}

// CreateBook func for creates a new book.
// @Summary Create a new book
// @Tags Book
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Success 200 {object} entities.Book
// @Security ApiKeyAuth
// @Router /book [post]
func CreateBook(c *fiber.Ctx) error {
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
}

// UpdateBook func for updates book by given ID.
// @Summary Update book
// @Tags Book
// @Accept json
// @Produce json
// @Param id body string true "Book ID"
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param book_status body integer true "Book status"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /book [put]
func UpdateBook(c *fiber.Ctx) error {
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
}

// DeleteBook func for deletes book by given ID.
// @Summary Delete book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id body string true "Book ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /book [delete]
func DeleteBook(c *fiber.Ctx) error {
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
}
