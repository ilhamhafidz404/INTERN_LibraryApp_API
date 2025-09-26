package controllers

import (
	"library_app/database"
	"library_app/dto"
	"library_app/helpers"
	"library_app/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Get Admins godoc
// @Summary Get all Admins
// @Description Ambil semua data admin
// @Tags Admins
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /api/admins [get]
func GetAdmins(c *fiber.Ctx) error {
	var admins []dto.Admin

	if err := database.DB.Model(&models.Admin{}).
		Order("id DESC").
		Where("level = ?", "admin").
		Scan(&admins).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed get data Admins")
	}

	return helpers.ResponseSuccess(c, "Success get data Admins", admins)
}

// StoreAdmin godoc
// @Summary Create Admin
// @Description Tambah data admin
// @Tags Admins
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param name formData string true "Nama Lengkap"
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Param confirmation_password formData string true "Konfirmasi Password"
// @Router /api/admins [post]
func StoreAdmin(c *fiber.Ctx) error {
	name := c.FormValue("name")
	username := c.FormValue("username")
	password := c.FormValue("password")
	confirmationPassword := c.FormValue("confirmation_password")

	if password != confirmationPassword {
		return helpers.ResponseError(c, "ALP-001", "Password dan konfirmasi tidak cocok")
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return helpers.ResponseError(c, "ALP-003", "Gagal enkripsi password")
	}

	admin := models.Admin{
		Name:     name,
		Username: username,
		Password: hashedPassword,
		Level:    "admin", // default level
	}

	if err := database.DB.Create(&admin).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal menyimpan data admin")
	}

	responseAdmin := dto.Admin{
		ID:       int(admin.ID),
		Name:     admin.Name,
		Username: admin.Username,
	}

	return helpers.ResponseSuccess(c, "Create Admin Success", responseAdmin)
}

// UpdateAdmin godoc
// @Summary Update Admin
// @Description Update data admin
// @Tags Admins
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param admin_id path int true "ID Admin"
// @Param name formData string false "Nama Lengkap"
// @Param username formData string false "Username"
// @Param password formData string false "Password"
// @Param confirmation_password formData string false "Konfirmasi Password"
// @Router /api/admins/{admin_id} [put]
func UpdateAdmin(c *fiber.Ctx) error {
	adminID := c.Params("admin_id")

	var admin models.Admin
	if err := database.DB.First(&admin, adminID).Error; err != nil {
		return helpers.ResponseError(c, "ALP-004", "Data admin tidak ditemukan")
	}

	name := c.FormValue("name")
	username := c.FormValue("username")
	password := c.FormValue("password")
	confirmationPassword := c.FormValue("confirmation_password")

	// update password jika ada
	if password != "" {
		if password != confirmationPassword {
			return helpers.ResponseError(c, "ALP-001", "Password dan konfirmasi tidak cocok")
		}
		hashedPassword, err := helpers.HashPassword(password)
		if err != nil {
			return helpers.ResponseError(c, "ALP-003", "Gagal enkripsi password")
		}
		admin.Password = hashedPassword
	}

	if name != "" {
		admin.Name = name
	}
	if username != "" {
		admin.Username = username
	}

	admin.UpdatedAt = time.Now()

	if err := database.DB.Save(&admin).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal mengupdate data admin")
	}

	responseAdmin := dto.Admin{
		ID:       int(admin.ID),
		Name:     admin.Name,
		Username: admin.Username,
	}

	return helpers.ResponseSuccess(c, "Update Admin Success", responseAdmin)
}

// DeleteAdmin godoc
// @Summary Delete Admin
// @Description Hapus Data Admin
// @Tags Admins
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param admin_id path int true "ID Admin"
// @Router /api/admins/{admin_id} [delete]
func DeleteAdmin(c *fiber.Ctx) error {
	adminID := c.Params("admin_id")
	if adminID == "" {
		return helpers.ResponseError(c, "ALP-002", "ID tidak ditemukan di URL")
	}

	var admin models.Admin
	if err := database.DB.First(&admin, adminID).Error; err != nil {
		return helpers.ResponseError(c, "ALP-002", "Admin tidak ditemukan")
	}

	if err := database.DB.Delete(&admin).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal menghapus admin")
	}

	return helpers.ResponseSuccess(c, "Delete Admin Success", nil)
}
