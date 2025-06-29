package routes

import (
	"library_app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    api.Get("/books", controllers.GetBooks)
}
