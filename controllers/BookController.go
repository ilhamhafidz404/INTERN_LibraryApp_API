package controllers

import (
	"library_app/database"
	"library_app/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
    var books []models.Book

    if err := database.DB.Find(&books).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error":   true,
            "message": "Failed to retrieve books",
        })
    }

    return c.JSON(fiber.Map{
        "error": false,
        "data":  books,
    })
}
