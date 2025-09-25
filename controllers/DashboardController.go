package controllers

import (
	"library_app/database"
	"library_app/dto"
	"library_app/helpers"
	"library_app/models"

	"github.com/gofiber/fiber/v2"
)

// GetDashboard godoc
// @Summary Get Dashboard Information
// @Description Ambil data dashboard
// @Tags Dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /api/dashboard [get]
func GetDashboard(c *fiber.Ctx) error {
	var dashboard dto.Dashboard

	if err := database.DB.Model(&models.Book{}).Count(&dashboard.TotalBook).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed get count Books")
	}

	if err := database.DB.Model(&models.Admin{}).Count(&dashboard.TotalAdmin).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed get count Admin")
	}

	if err := database.DB.Model(&models.Student{}).Count(&dashboard.TotalStudent).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed get count Student")
	}

	return helpers.ResponseSuccess(c, "Success get data Dashboard", dashboard)
}