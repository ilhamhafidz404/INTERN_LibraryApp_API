package controllers

import (
	"library_app/database"
	"library_app/dto"
	"library_app/helpers"
	"library_app/models"

	_ "library_app/dto"

	"github.com/gofiber/fiber/v2"
)

// GetStudents godoc
// @Summary Get all students
// @Description Ambil semua data buku
// @Tags Students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /api/students [get]
func GetStudents(c *fiber.Ctx) error {
    var students []dto.Student

	if err := database.DB.Model(&models.Student{}).
		Scan(&students).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed get data Students")
	}

	return helpers.ResponseSuccess(c, "Success get data Students", students)
}