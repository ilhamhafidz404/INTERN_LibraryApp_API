package routes

import (
	"library_app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

	// Auth
    api.Post("/login", controllers.Login)
    api.Post("/register", controllers.Register)

	// Books
    api.Get("/books", controllers.GetBooks)
}
