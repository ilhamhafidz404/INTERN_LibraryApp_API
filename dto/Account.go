package dto

import (
	"time"
)

type Student struct {
	NISN         string    `json:"nisn"`
	NIK          string    `json:"nik"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	PlaceOfBirth string    `json:"place_of_birth"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	MotherName   string    `json:"mother_name"`
	Gender       string    `json:"gender"`
	Level        string    `json:"level"`
}

type StudentUpdateProfileRequest struct {
	NISN         string    `json:"nisn"`
	NIK          string    `json:"nik"`
	Name         string    `json:"name"`
	PlaceOfBirth string    `json:"place_of_birth"`
	DateOfBirth  string    `json:"date_of_birth"`
	MotherName   string    `json:"mother_name"`
	Gender       string    `json:"gender"`
	Level        string    `json:"level"`
}

type StudentChangePasswordRequest struct {
	OldPassword         		string    `json:"old_password"`
	NewPassword         		string    `json:"new_password"`
	ConfirmationNewPassword     string    `json:"confirmation_new_password"`
}