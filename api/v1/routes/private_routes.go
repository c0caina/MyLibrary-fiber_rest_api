package routes

import (
	"time"

	"github.com/c0caina/MyLibrary-fiber_rest_api/database/postgres"
	"github.com/c0caina/MyLibrary-fiber_rest_api/entities"
	"github.com/c0caina/MyLibrary-fiber_rest_api/pkg/middleware"
	utils "github.com/c0caina/MyLibrary-fiber_rest_api/pkg/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtected(), func (c *fiber.Ctx) error {

		claims, err := utils.ParceJWT(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if time.Now().Unix() > claims.Expires {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized, check expiration time of your token",
			})
		}
	
		book := &entities.Book{}
		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		db, err := postgres.NewPostgres()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
	
		book.ID = uuid.New()
		book.CreatedAt = time.Now()
		book.BookStatus = 1
	
		if err := utils.NewValidator().Struct(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ValidatorErrors(err),
			})
		}
	
		if err := db.CreateBook(book); err != nil {
	
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		return c.JSON(fiber.Map{
			"error": err,
			"book":  book,
		})
	})
	route.Put("/book", middleware.JWTProtected(), func (c *fiber.Ctx) error {

		claims, err := utils.ParceJWT(c)
		if err != nil {
	
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if time.Now().Unix() > claims.Expires {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized, check expiration time of your token",
			})
		}
	
		book := &entities.Book{}
		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		db, err := postgres.NewPostgres()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		book.UpdatedAt = time.Now()
		if err := utils.NewValidator().Struct(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ValidatorErrors(err),
			})
		}
	
		if err := db.UpdateBook(book); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		return c.SendStatus(fiber.StatusCreated)
	})
	route.Delete("/book", middleware.JWTProtected(), func (c *fiber.Ctx) error {

		claims, err := utils.ParceJWT(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if time.Now().Unix() > claims.Expires {
	
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized, check expiration time of your token",
			})
		}
	
		book := &entities.Book{}
		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		if err := utils.NewValidator().StructPartial(book, "id"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ValidatorErrors(err),
			})
		}
	
		db, err := postgres.NewPostgres()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		foundedBook, err := db.GetBook(book.ID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "book with this ID not found",
			})
		}
	
		if err := db.DeleteBook(foundedBook.ID); err != nil {
	
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	
		return c.SendStatus(fiber.StatusNoContent)
	})
}
