package main

import (
	"log"
	
	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New()

	if err := app.Listen(":3000"); err !=nil{
		log.Printf("Server is not running! Reason: %v", err)
	}
}