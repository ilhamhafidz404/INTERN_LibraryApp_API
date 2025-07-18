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
	var response dto.LoginResponse

	// Parse dan validasi body
	if err := c.BodyParser(&payload); err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid request body")
	}

	// Validasi struktur payload
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return helpers.ResponseError(c, "ALP-003", "Validation Failed: "+err.Error())
	}

	// Coba cari siswa terlebih dahulu
	var student models.Student
	result := database.DB.Where("nisn = ?", payload.Username).First(&student)

	if result.Error == nil && result.RowsAffected > 0 {
		// Password cocok?
		if !helpers.CheckPassword(payload.Password, student.Password) {
			return helpers.ResponseError(c, "ALP-003", "Invalid Password")
		}

		// Generate JWT
		token, err := helpers.GenerateJWT(student.ID, student.NISN)
		if err != nil {
			return helpers.ResponseError(c, "ALP-004", "Gagal generate token")
		}

		// Buat response
		loginUserResponse := dto.LoginUserResponse{
			ID:                  student.ID,
			Name:                student.Name,
			Username:            student.NISN,
			// Level:               student.Level,
			Level:               "student",
		}

		response = dto.LoginResponse{
			Token: token,
			Auth:  loginUserResponse,
		}

		return helpers.ResponseSuccess(c, "Login Success", response)
	}

	// Jika bukan siswa, coba cari admin
	var admin models.Admin
	result = database.DB.Where("username = ?", payload.Username).First(&admin)

	if result.Error == nil && result.RowsAffected > 0 {
		// Validasi password admin
		if !helpers.CheckPassword(payload.Password, admin.Password) {
			return helpers.ResponseError(c, "ALP-004", "Invalid Password")
		}

		// Buat token admin
		token, err := helpers.GenerateJWT(admin.ID, admin.Username)
		if err != nil {
			return helpers.ResponseError(c, "ALP-004", "Gagal generate token")
		}

		loginUserResponse := dto.LoginUserResponse{
			ID:                  admin.ID,
			Name:                admin.Name,
			Username:            admin.Username,
			Level:               admin.Level,
		}

		response = dto.LoginResponse{
			Token: token,
			Auth:  loginUserResponse,
		}

		return helpers.ResponseSuccess(c, "Login Success", response)
	}

	// Jika tidak ditemukan di student maupun admin
	return helpers.ResponseError(c, "ALP-002", "User not found")

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

