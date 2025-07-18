package dto

type Book struct {
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Year      uint16 `json:"year"`
	Total     int64  `json:"total"`
	CreatedBy uint   `json:"creted_by"`
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
