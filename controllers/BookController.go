package controllers

import (
	"library_app/database"
	"library_app/dto"
	"library_app/helpers"
	"library_app/models"

	_ "library_app/dto"

	"github.com/go-playground/validator/v10"
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

    // 1. Parse body
	if err := c.BodyParser(&payload); err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Request Body")
	}

    // 2. Validasi
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return helpers.ResponseError(c, "ALP-003", "Validation Failed: "+err.Error())
	}

    // 3. Insert Book
	book := models.Book{
		Title: payload.Title,
        Publisher: payload.Publisher,
        Author: payload.Author,
        ISBN: payload.ISBN,
        Year: payload.Year,
        Total: payload.Total,
        CreatedBy: 1,
	}

	// 4. Simpan ke database
	if err := database.DB.Create(&book).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed Insert Book")
	}

    return helpers.ResponseSuccess(c, "Create Success", book)
}

// PutBooks godoc
// @Summary Update Book
// @Description Update data buku
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.ResponseSuccess
// @Failure 500 {object} dto.ResponseError
// @Param request body dto.BookRequest true "Book payload"
// @Param id path int true "ID Buku"
// @Router /api/books/{id} [put]
func UpdateBook(c *fiber.Ctx) error {
    var payload dto.BookRequest

    // 1. Ambil ID dari parameter
	id := c.Params("id")
	if id == "" {
		return helpers.ResponseError(c, "ALP-001", "ID tidak ditemukan di URL")
	}

	// 2. Parsing body
	if err := c.BodyParser(&payload); err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Request Body")
	}

	// 3. Validasi payload
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return helpers.ResponseError(c, "ALP-003", "Validation Failed: "+err.Error())
	}

	// 4. Cek apakah buku ada
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return helpers.ResponseError(c, "ALP-002", "Buku tidak ditemukan")
	}

	// 5. Update field dari payload
	book.Title = payload.Title
	book.Publisher = payload.Publisher
	book.Author = payload.Author
	book.ISBN = payload.ISBN
	book.Year = payload.Year
	book.Total = payload.Total
    book.CreatedBy = 1

	// 6. Simpan perubahan
	if err := database.DB.Save(&book).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal update buku")
	}

	return helpers.ResponseSuccess(c, "Update Success", book)
}
