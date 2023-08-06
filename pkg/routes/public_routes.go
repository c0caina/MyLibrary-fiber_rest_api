package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/c0caina/MyLibrary-fiber_rest_api/internal/app/controllers"
)


func PublicRoutes(a *fiber.App) {
	
	route := a.Group("/api/v1")

	route.Get("/books", controllers.GetBooks)              
	route.Get("/book/:id", controllers.GetBook)            
	route.Get("/token/new", controllers.GetNewAccessToken) 
}
