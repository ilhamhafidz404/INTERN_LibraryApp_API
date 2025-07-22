package dto

type LendingHistory struct {
	ID            uint   `json:"id"`
	BookID        uint   `json:"book_id"`
	BookTitle     string `json:"book_title"`
	BookAuthor    string `json:"book_author"`
	BookPublisher string `json:"book_publisher"`
	BookISBN      string `json:"book_isbn"`
	StudentID     uint   `json:"student_id"`
	StudentNISN   string `json:"student_nisn"`
	StudentName   string `json:"student_name"`
	Status        string `json:"status"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}