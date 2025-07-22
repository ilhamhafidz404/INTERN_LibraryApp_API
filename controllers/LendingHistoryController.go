package controllers

import (
	"library_app/database"
	"library_app/dto"
	"library_app/helpers"

	_ "library_app/dto"

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
			1=1
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

	return helpers.ResponseSuccess(c, "Berhasil mengambil data peminjaman", histories)
}

