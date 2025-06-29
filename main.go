package main

import (
	"library_app/database"
	"library_app/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // migration
    database.Connect()
    database.DB.AutoMigrate(
        models.Book{},
        &models.User{},
    )

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Fiber!")
    })

    app.Listen(":3000")
}
