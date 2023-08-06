package routes

import (
	"github.com/c0caina/MyLibrary-fiber_rest_api/database/postgres"
	utils "github.com/c0caina/MyLibrary-fiber_rest_api/pkg/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PublicRoutes(a *fiber.App) {

	route := a.Group("/api/v1")

	route.Get("/books", func (c *fiber.Ctx) error {

		db, err := postgres.NewPostgres()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		books, err := db.GetBooks()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		return c.JSON(fiber.Map{
			"error": err,
			"count": len(books),
			"books": books,
		})
	})
	route.Get("/book/:id", func (c *fiber.Ctx) error {

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		db, err := postgres.NewPostgres()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		book, err := db.GetBook(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
				"book":  book,
			})
		}
	
		return c.JSON(fiber.Map{
			"error": err,
			"book":  book,
		})
	})
	route.Get("/token/new", func (c *fiber.Ctx) error {

		token, err := utils.GenerateAccessJWT()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		return c.JSON(fiber.Map{
			"error":        err,
			"access_token": token,
		})
	})
}
