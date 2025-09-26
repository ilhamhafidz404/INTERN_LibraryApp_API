package controllers

import (
	"fmt"
	"library_app/database"
	"library_app/dto"
	"library_app/helpers"
	"library_app/models"
	"time"

	_ "library_app/dto"

	"github.com/gofiber/fiber/v2"
)

// GetStudents godoc
// @Summary Get all students
// @Description Ambil semua data Student
// @Tags Students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /api/students [get]
func GetStudents(c *fiber.Ctx) error {
    var students []dto.Student

	if err := database.DB.Model(&models.Student{}).Order("id DESC").
		Scan(&students).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed get data Students")
	}

	return helpers.ResponseSuccess(c, "Success get data Students", students)
}

// StoreStudent godoc
// @Summary Create Student
// @Description Create data siswa dengan upload foto opsional
// @Tags Students
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param nisn formData string true "NISN"
// @Param nik formData string true "NIK"
// @Param name formData string true "Nama Lengkap"
// @Param password formData string true "Password"
// @Param confirmation_password formData string true "Konfirmasi Password"
// @Param place_of_birth formData string true "Tempat Lahir"
// @Param date_of_birth formData string true "Tanggal Lahir (YYYY-MM-DD)"
// @Param mother_name formData string true "Nama Ibu"
// @Param gender formData string true "Jenis Kelamin (M/F)"
// @Param level formData string true "Tingkat Kelas"
// @Router /api/students [post]
func StoreStudent(c *fiber.Ctx) error {
	// 1. Ambil form data
	nisn := c.FormValue("nisn")
	nik := c.FormValue("nik")
	name := c.FormValue("name")
	password := c.FormValue("password")
	confirmationPassword := c.FormValue("confirmation_password")
	placeOfBirth := c.FormValue("place_of_birth")
	dateOfBirth := c.FormValue("date_of_birth")
	motherName := c.FormValue("mother_name")
	gender := c.FormValue("gender")
	level := c.FormValue("level")

	// Validasi password dan konfirmasi
	if password != confirmationPassword {
		return helpers.ResponseError(c, "ALP-001", "Password dan konfirmasi tidak cocok")
	}

	// 2. Handle upload file (opsional)
	// photoFile, err := c.FormFile("photo")
	// var photoFilename string
	// if err == nil && photoFile != nil {
	// 	photoFilename = fmt.Sprintf("student_%d_%s", time.Now().UnixNano(), photoFile.Filename)
	// 	savePath := fmt.Sprintf("./uploads/%s", photoFilename)

	// 	if err := c.SaveFile(photoFile, savePath); err != nil {
	// 		log.Printf("Gagal menyimpan file foto: %v", err)
	// 		return helpers.ResponseError(c, "ALP-006", "Gagal upload file foto")
	// 	}
	// }

	// 3. Parse date_of_birth
	var err error

	var dob time.Time
	if dateOfBirth != "" {
		dob, err = time.Parse("2006-01-02", dateOfBirth)
		if err != nil {
			return helpers.ResponseError(c, "ALP-002", "Format tanggal lahir salah (YYYY-MM-DD)")
		}
	}

	// 4. Hash password
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return helpers.ResponseError(c, "ALP-003", "Gagal enkripsi password")
	}

	// 5. Buat data student
	student := models.Student{
		NISN:         nisn,
		NIK:          nik,
		Name:         name,
		Password:     hashedPassword,
		PlaceOfBirth: placeOfBirth,
		DateOfBirth:  dob,
		MotherName:   motherName,
		Gender:       gender,
		Level:        level,
	}

	// 6. Simpan ke database
	if err := database.DB.Create(&student).Error; err != nil {
		fmt.Println(err);
		return helpers.ResponseError(c, "ALP-005", "Gagal menyimpan data siswa")
	}

	// 7. Response
	responseStudent := dto.Student{
		ID:           int(student.ID),
		NISN:         student.NISN,
		NIK:          student.NIK,
		Name:         student.Name,
		PlaceOfBirth: student.PlaceOfBirth,
		DateOfBirth:  student.DateOfBirth,
		MotherName:   student.MotherName,
		Gender:       student.Gender,
		Level:        student.Level,
	}

	return helpers.ResponseSuccess(c, "Create Success", responseStudent)
}

// UpdateStudent godoc
// @Summary Update Student
// @Description Update data siswa dengan upload foto opsional
// @Tags Students
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param student_id path int true "ID Student"
// @Param nisn formData string false "NISN"
// @Param nik formData string false "NIK"
// @Param name formData string false "Nama Lengkap"
// @Param password formData string false "Password (opsional)"
// @Param confirmation_password formData string false "Konfirmasi Password"
// @Param place_of_birth formData string false "Tempat Lahir"
// @Param date_of_birth formData string false "Tanggal Lahir (YYYY-MM-DD)"
// @Param mother_name formData string false "Nama Ibu"
// @Param gender formData string false "Jenis Kelamin (M/F)"
// @Param level formData string false "Tingkat Kelas"
// @Router /api/students/{id} [put]
func UpdateStudent(c *fiber.Ctx) error {
	// 1. Ambil ID student dari path
	student_id := c.Params("student_id")

	// 2. Cari student di DB
	var student models.Student
	if err := database.DB.First(&student, student_id).Error; err != nil {
		return helpers.ResponseError(c, "ALP-004", "Data siswa tidak ditemukan")
	}

	fmt.Print("id", student_id);

	// 3. Ambil form data
	nisn := c.FormValue("nisn")
	nik := c.FormValue("nik")
	name := c.FormValue("name")
	password := c.FormValue("password")
	confirmationPassword := c.FormValue("confirmation_password")
	placeOfBirth := c.FormValue("place_of_birth")
	dateOfBirth := c.FormValue("date_of_birth")
	motherName := c.FormValue("mother_name")
	gender := c.FormValue("gender")
	level := c.FormValue("level")

	// 4. Validasi password (kalau ada)
	if password != "" {
		if password != confirmationPassword {
			return helpers.ResponseError(c, "ALP-001", "Password dan konfirmasi tidak cocok")
		}
		hashedPassword, err := helpers.HashPassword(password)
		if err != nil {
			return helpers.ResponseError(c, "ALP-003", "Gagal enkripsi password")
		}
		student.Password = hashedPassword
	}

	// 5. Parse tanggal lahir (kalau ada)
	if dateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", dateOfBirth)
		if err != nil {
			return helpers.ResponseError(c, "ALP-002", "Format tanggal lahir salah (YYYY-MM-DD)")
		}
		student.DateOfBirth = dob
	}

	// 6. Update field lain (kalau tidak kosong)
	if nisn != "" {
		student.NISN = nisn
	}
	if nik != "" {
		student.NIK = nik
	}
	if name != "" {
		student.Name = name
	}
	if placeOfBirth != "" {
		student.PlaceOfBirth = placeOfBirth
	}
	if motherName != "" {
		student.MotherName = motherName
	}
	if gender != "" {
		student.Gender = gender
	}
	if level != "" {
		student.Level = level
	}

	// 7. Simpan update ke database
	if err := database.DB.Save(&student).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal mengupdate data siswa")
	}

	// 8. Response
	responseStudent := dto.Student{
		ID:           int(student.ID),
		NISN:         student.NISN,
		NIK:          student.NIK,
		Name:         student.Name,
		PlaceOfBirth: student.PlaceOfBirth,
		DateOfBirth:  student.DateOfBirth,
		MotherName:   student.MotherName,
		Gender:       student.Gender,
		Level:        student.Level,
	}

	return helpers.ResponseSuccess(c, "Update Success", responseStudent)
}



// DELETE Student godoc
// @Summary Delete Student
// @Description Hapus Data Siswa
// @Tags Students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param student_id path int true "ID Student"
// @Router /api/students/{student_id} [delete]
func DeleteStudent(c *fiber.Ctx) error {
	// 1. Ambil ID dari parameter
	studentID := c.Params("student_id")
	if studentID == "" {
		return helpers.ResponseError(c, "ALP-002", "ID tidak ditemukan di URL")
	}

	// 2. Cek apakah student ada
	var student models.Student
	if err := database.DB.First(&student, studentID).Error; err != nil {
		return helpers.ResponseError(c, "ALP-002", "siswa tidak ditemukan")
	}

	// 3. Hapus student
	if err := database.DB.Delete(&student).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal menghapus siswa")
	}

	return helpers.ResponseSuccess(c, "Delete Siswa Success", nil)
}