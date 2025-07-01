package dto

// ALP-001 = SUCCESS
// ALP-002 = Retun 404
// ALP-003 = Validasi
// ALP-004 = Error
// ALP-005 = Error Database

type ResponseSuccess struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Success bool     `json:"success"`
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}
