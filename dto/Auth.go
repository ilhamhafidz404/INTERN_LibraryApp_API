package dto

// LOGIN
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	ID    					uint   	`json:"id"`
	Name  					string 	`json:"name"`
	NISN  					string 	`json:"nisn"`
	NIK   					string	`json:"nik"` 
	PlaceAndDateOfBirth 	string 	`json:"place_and_date_of_birth"`       
    MotherName   			string 	`json:"mother_name"`           
    Gender       			string 	`json:"gender"`
    Level 		 			string 	`json:"level"` 
}

// REGISTER
type RegisterRequest struct {
	NISN         	string 	`json:"nisn" validate:"min=10,max=10"`         
    NIK          	string 	`json:"nik" validate:"max=16"`  
    Name         	string 	`json:"name" validate:"required"`              
    Password        string 	`json:"password" validate:"required"`
	ConfirmPassword string 	`json:"confirm_password" validate:"required"`     
    PlaceOfBirth 	string 	`json:"place_of_birth" validate:"required"`    
    DateOfBirth  	string	`json:"date_of_birth" validate:"required"`       
    MotherName   	string 	`json:"mother_name"`           
    Gender       	string 	`json:"gender" validate:"required,oneof=M F"`
    Level 		 	string 	`json:"level" validate:"required,oneof=X XI XII"`
}

type RegisterResponse struct {
	ID    uint   	`json:"id"`
	Name  string 	`json:"name"`
	NISN  string 	`json:"nisn"`
	NIK   string	`json:"nik"`  
}
