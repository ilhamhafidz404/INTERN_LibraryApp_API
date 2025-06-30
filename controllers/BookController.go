package controllers

import (
	"library_app/database"
	"library_app/helpers"
	"library_app/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
    var books []models.Book

    if err := database.DB.Find(&books).Error; err != nil {
        return helpers.ResponseError(c, "ALP-003", "Failed get data Books")
    }

    return helpers.ResponseSuccess(c, "Success get data Books", books)
}
