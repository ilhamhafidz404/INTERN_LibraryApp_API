package controllers

import (
	"library_app/database"
	"library_app/dto"
	"library_app/helpers"
	"library_app/models"

	_ "library_app/dto"

	"github.com/gofiber/fiber/v2"
)

// GetBooks godoc
// @Summary Get all books
// @Description Ambil semua data buku
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.ResponseSuccess
// @Failure 500 {object} dto.ResponseError
// @Router /api/books [get]
func GetBooks(c *fiber.Ctx) error {
    var books []models.Book

    if err := database.DB.Find(&books).Error; err != nil {
        return helpers.ResponseError(c, "ALP-003", "Failed get data Books")
    }

    return helpers.ResponseSuccess(c, "Success get data Books", books)
}

// PostBooks godoc
// @Summary Create Book
// @Description Create data buku
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.ResponseSuccess
// @Failure 500 {object} dto.ResponseError
// @Param request body dto.BookRequest true "Book payload"
// @Router /api/books [post]
func StoreBook(c *fiber.Ctx) error {
    var payload dto.BookRequest

    return helpers.ResponseSuccess(c, "Success get data Books", payload)
}
