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

// Get Profile Detail godoc
// @Summary Get Student Profile
// @Description Lihat Detail Profile Siswa
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param student_id path int true "Student id"
// @Router /api/profile/{student_id} [get]
func GetProfile(c *fiber.Ctx) error {
    var account dto.Student

	studentID := c.Params("student_id")
	if studentID == "" {
		return helpers.ResponseError(c, "ALP-002", "ID tidak ditemukan di URL")
	}

	if err := database.DB.Model(&models.Student{}).
		First(&account, studentID).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Error Get Database")
	}

	return helpers.ResponseSuccess(c, "Success get data Student", account)
}

// Put Profile godoc
// @Summary Update Profile
// @Description Update Student Profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.StudentUpdateProfileRequest true "Student Request payload"
// @Param student_id path int true "Student id"
// @Router /api/profile/{student_id} [put]
func UpdateProfile(c *fiber.Ctx) error {
    var payload dto.StudentUpdateProfileRequest

    // 1. Ambil ID dari parameter
	studentID := c.Params("student_id")
	if studentID == "" {
		return helpers.ResponseError(c, "ALP-002", "ID tidak ditemukan di URL")
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

	// 4. Cek apakah student ada
	var student models.Student
	if err := database.DB.First(&student, studentID).Error; err != nil {
		return helpers.ResponseError(c, "ALP-002", "Siswa tidak ditemukan")
	}

	// 5. Parse ke Time.time
	parsedDateOfBirth, err := time.Parse("02-01-2006", payload.DateOfBirth)
	if err != nil {
		return helpers.ResponseError(c, "ALP-004", "Format tanggal salah")
	}

	// 6. Update field dari payload
	student.DateOfBirth= parsedDateOfBirth
	student.Gender= payload.Gender
	student.Level = payload.Level
	student.MotherName = payload.MotherName
	student.NIK = payload.NIK
	student.NISN = payload.NISN
	student.Name = payload.Name
	student.PlaceOfBirth = payload.PlaceOfBirth

	// 7. Simpan perubahan
	if err := database.DB.Save(&student).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal update profile")
	}

	// 8. Preparation Response
	responseProfile := dto.Student {
		NISN: student.NISN,
		NIK: student.NIK,
		Name: student.Name,
		PlaceOfBirth: student.PlaceOfBirth,
		DateOfBirth: student.DateOfBirth,
		MotherName: student.MotherName,
		Gender: student.Gender,
		Level: student.Level,
	}

	return helpers.ResponseSuccess(c, "Update Success", responseProfile)
}

// Put Profile godoc
// @Summary Update Profile
// @Description Update Student Profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.StudentChangePasswordRequest true "Student Change Password Request payload"
// @Param student_id path int true "Student id"
// @Router /api/profile/change-password/{student_id} [put]
func ChangePassword(c *fiber.Ctx) error {
    var payload dto.StudentChangePasswordRequest

    // 1. Ambil ID dari parameter
	studentID := c.Params("student_id")
	if studentID == "" {
		return helpers.ResponseError(c, "ALP-002", "ID tidak ditemukan di URL")
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

	// 4. Cek apakah student sesuai dengan request
	var student models.Student
	if err := database.DB.First(&student, studentID).Error; err != nil {
		return helpers.ResponseError(c, "ALP-002", "Siswa tidak ditemukan")
	}

	if !helpers.CheckPassword(payload.OldPassword, student.Password) {
		return helpers.ResponseError(c, "ALP-003", "Invalid Old Password")
	}

	// 5. Validasi Password dan Confirm Password
	if payload.NewPassword != payload.ConfirmationNewPassword {
		return helpers.ResponseError(c, "ALP-003", "Validation Failed: Password Mismatch")
	}

	// 6. Hash password
	hashedPassword, err := helpers.HashPassword(payload.NewPassword)
	if err != nil {
		return helpers.ResponseError(c, "ALP-004", "Failed Encrypt Password")
	}

	// 7. Update field dari payload
	student.Password= hashedPassword

	// 8. Simpan perubahan
	if err := database.DB.Save(&student).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal update profile")
	}

	// 9. Preparation Response
	responseProfile := dto.Student {
		NISN: student.NISN,
		NIK: student.NIK,
		Name: student.Name,
		PlaceOfBirth: student.PlaceOfBirth,
		DateOfBirth: student.DateOfBirth,
		MotherName: student.MotherName,
		Gender: student.Gender,
		Level: student.Level,
	}

	return helpers.ResponseSuccess(c, "Change Password Success", responseProfile)
}
