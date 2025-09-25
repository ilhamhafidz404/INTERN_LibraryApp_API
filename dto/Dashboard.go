package dto

type Dashboard struct {
	TotalAdmin   int64 `json:"total_admin"`
	TotalStudent int64 `json:"total_student"`
	TotalBook    int64 `json:"total_book"`
}