package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New()

	if err := app.Listen(os.Getenv("SERVER_URL")); err !=nil{
		log.Printf("Server is not running! Reason: %v", err)
	}
}