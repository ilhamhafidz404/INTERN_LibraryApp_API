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

// Get Lending History godoc
// @Summary Get Lending History
// @Description Lihat History Peminjaman
// @Tags Lending History
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param student_id query int false "Filter berdasarkan ID siswa"
// @Param start_date query string false "Filter tanggal mulai"
// @Param end_date query string false "Filter tanggal selesai"
// @Router /api/lending-history [get]
func GetLendingHistory(c *fiber.Ctx) error {
	var histories []dto.LendingHistory

	studentID := c.Query("student_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	// Bangun SQL query secara dinamis
	sqlQuery := `
		SELECT 
			lh.id, 
			lh.created_at, 
			lh.updated_at, 
			lh.status,
			lh.start_date,
			lh.end_date,
			s.id AS student_id, 
			s.nisn AS student_nisn, 
			s.name AS student_name,
			b.id AS book_id, 
			b.title AS book_title, 
			b.author AS book_author,
			b.publisher AS book_publisher, 
			b.isbn AS book_isbn
		FROM 
			lending_histories lh
		JOIN 
			students s ON s.id = lh.student_id
		JOIN 
			books b ON b.id = lh.book_id
		WHERE
			lh.deleted_at IS NULL
	`

	// Tambahkan filter jika ada
	var args []interface{}

	if studentID != "" {
		sqlQuery += " AND lh.student_id = ?"
		args = append(args, studentID)
	}

	if startDate != "" && endDate != "" {
		sqlQuery += " AND lh.created_at BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	if err := database.DB.Raw(sqlQuery, args...).Scan(&histories).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal mengambil data history")
	}

	if len(histories) <= 0 {
		return helpers.ResponseError(c, "ALP-002", "History Peminjaman Kosong")
	}
	
	return helpers.ResponseSuccess(c, "Berhasil mengambil data peminjaman", histories)
}

// POST Lending History godoc
// @Summary Post Lending History
// @Description Tambah Data History Peminjaman
// @Tags Lending History
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.LendingHistoryRequest true "Landing History payload"
// @Router /api/lending-history [post]
func PostLendingHistory(c *fiber.Ctx) error {
	var payload dto.LendingHistoryRequest

	// 1. Parse body
	if err := c.BodyParser(&payload); err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Request Body")
	}

	// 2. Validasi
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return helpers.ResponseError(c, "ALP-003", "Validation Failed: "+err.Error())
	}

	// 3. Convert Date
	StartDate, err := time.Parse("2006-01-02", payload.StartDate)
	if err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Date of Start Date (use YYYY-MM-DD)")
	}

	EndDate, err := time.Parse("2006-01-02", payload.EndDate)
	if err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Date of End Date (use YYYY-MM-DD)")
	}

	// 4. Insert User
	history := models.LendingHistory{
		BookID: payload.BookID,
		StudentID: payload.StudentID,
		StartDate: StartDate,
		EndDate: EndDate,
		Status: "loaned",
	}

	// 5. Simpan ke database
	if err := database.DB.Create(&history).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed Insert Lending History")
	}

	// 5. Prepare for Response
	response := dto.LendingHistoryResponse{
		StartDate: payload.StartDate,
		EndDate: payload.EndDate,
	}

	//
	return helpers.ResponseSuccess(c, "Create History Success", response)
}

// PUT Lending History godoc
// @Summary Put Lending History
// @Description Perbarui Data History Peminjaman
// @Tags Lending History
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param history_id path int true "ID Landing History"
// @Param request body dto.LendingHistoryRequest true "Landing History payload"
// @Router /api/lending-history/{history_id} [put]
func PutLendingHistory(c *fiber.Ctx) error {
	var payload dto.LendingHistoryRequest
	
	// 1. Parse body
	if err := c.BodyParser(&payload); err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Request Body")
	}

	// 2. Ambil ID dari parameter
	historyID := c.Params("history_id")
	if historyID == "" {
		return helpers.ResponseError(c, "ALP-002", "ID tidak ditemukan di URL")
	}

	// 3. Validasi
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return helpers.ResponseError(c, "ALP-003", "Validation Failed: "+err.Error())
	}

	// 4. Cek apakah history ada
	var lendingHistory models.LendingHistory
	if err := database.DB.First(&lendingHistory, historyID).Error; err != nil {
		return helpers.ResponseError(c, "ALP-002", "History tidak ditemukan")
	}

	// 5. Convert Date
	StartDate, err := time.Parse("2006-01-02", payload.StartDate)
	if err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Date of Start Date (use YYYY-MM-DD)")
	}

	EndDate, err := time.Parse("2006-01-02", payload.EndDate)
	if err != nil {
		return helpers.ResponseError(c, "ALP-004", "Invalid Date of End Date (use YYYY-MM-DD)")
	}

	// 6. Update Lending History
	lendingHistory.BookID = payload.BookID
	lendingHistory.StudentID = payload.StudentID
	lendingHistory.StartDate = StartDate
	lendingHistory.EndDate = EndDate
	lendingHistory.Status = payload.Status

	// 7. Simpan ke database
	if err := database.DB.Save(&lendingHistory).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Failed Update Lending History")
	}

	// 8. Prepare for Response
	response := dto.LendingHistoryResponse{
		StartDate: payload.StartDate,
		EndDate: payload.EndDate,
	}

	//
	return helpers.ResponseSuccess(c, "Update History Success", response)
}

// DELETE Lending History godoc
// @Summary Delete Lending History
// @Description Hapus Data History Peminjaman
// @Tags Lending History
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param history_id path int true "ID Landing History"
// @Router /api/lending-history/{history_id} [delete]
func DeleteLendingHistory(c *fiber.Ctx) error {
	// 1. Ambil ID dari parameter
	historyID := c.Params("history_id")
	if historyID == "" {
		return helpers.ResponseError(c, "ALP-002", "ID tidak ditemukan di URL")
	}

	// 2. Cek apakah history ada
	var lendingHistory models.LendingHistory
	if err := database.DB.First(&lendingHistory, historyID).Error; err != nil {
		return helpers.ResponseError(c, "ALP-002", "History tidak ditemukan")
	}

	// 3. Hapus history
	if err := database.DB.Delete(&lendingHistory).Error; err != nil {
		return helpers.ResponseError(c, "ALP-005", "Gagal menghapus buku")
	}

	return helpers.ResponseSuccess(c, "Delete History Success", nil)
}

