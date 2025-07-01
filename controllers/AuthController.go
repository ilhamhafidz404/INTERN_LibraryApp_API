package controllers

import (
	"library_app/database"
	"library_app/dto"
	"library_app/helpers"
	"library_app/models"
	"time"

	_ "library_app/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// AuthLogin godoc
// @Summary Post Login
// @Description Login Library App
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login payload"
// @Success 200 {object} dto.ResponseSuccess
// @Failure 400 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /api/login [post]
func Login(c *fiber.Ctx) error {
    var payload dto.LoginRequest

    // Parse dan validasi body
    if err := c.BodyParser(&payload); err != nil {
        return helpers.ResponseError(c, "ALP-004", "Invalid request body")
    }

    // Validasi field kosong
    if payload.Username == "" || payload.Password == "" {
        return helpers.ResponseError(c, "ALP-003", "Username dan password wajib diisi")
    }

	// Ambil data ke database
	// sqlQuery := `
	// 	SELECT
	// 		*
	// 	FROM
	// 		users
	// 	WHERE
	// 		nisn = ?
	// 	AND
	// 		password = ?
	// `

	return nil
}

// AuthRegister godoc
// @Summary Post Register
// @Description Register Library App
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register payload"
// @Success 200 {object} dto.ResponseSuccess
// @Failure 400 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /api/register [post]
func Register(c *fiber.Ctx) error {
	var payload dto.RegisterRequest

	// 1. Parse body
	if err := c.BodyParser(&payload); err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Request Body")
	}

	// 2. Validasi
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return helpers.ResponseError(c, "ALP-003", "Validation Failed: "+err.Error())
	}

	// 3. Validasi Password dan Confirm Password
	if payload.Password != payload.ConfirmPassword {
		return helpers.ResponseError(c, "ALP-003", "Validation Failed: Password Mismatch")
	}

	// 4. Hash password
	hashedPassword, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return helpers.ResponseError(c, "ALP-004", "Failed Encrypt Password")
	}

	// 5. Convert Date of Birth
	DateOfBirth, err := time.Parse("2006-01-02", payload.DateOfBirth)
	if err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Date of Birth Format (use YYYY-MM-DD)")
	}

	// 6. Insert User
	user := models.Student{
		NISN:         payload.NISN,
		NIK:          payload.NIK,
		Name:         payload.Name,
		Password:     hashedPassword,
		PlaceOfBirth: payload.PlaceOfBirth,
		DateOfBirth:  DateOfBirth,
		MotherName:   payload.MotherName,
		Gender:       payload.Gender,
		Level:        payload.Level,
	}

	// 7. Simpan ke database
	if err := database.DB.Create(&user).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed Insert User")
	}

	// 8. Prepare for Response
	response := dto.RegisterResponse{
		ID:    	user.ID,
		Name:  	user.Name,
		NISN:  	user.NISN,
		NIK: 	user.NIK,
	}

	//
	return helpers.ResponseSuccess(c, "Registration Success", response)
}

