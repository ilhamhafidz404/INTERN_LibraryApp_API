package dto

type Book struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	// Slug      string `json:"slug"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Year      uint16 `json:"year"`
	Total     int64  `json:"total"`
	CreatedBy uint   `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BookRequest struct {
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Year      uint16 `json:"year"`
	Total     int64  `json:"total"`
}
