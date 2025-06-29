package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    NISN         string 	`gorm:"type:varchar(10);unique;not null" json:"nisn"`          
    NIK          string 	`gorm:"type:varchar(16);unique;not null" json:"nik"`          
    Name         string 	`gorm:"type:varchar(50);not null" json:"name"`               
    PlaceOfBirth string 	`gorm:"type:varchar(30);not null" json:"place_of_birth"`      
    DateOfBirth  time.Time 	`gorm:"type:datetime;not null" json:"date_of_birth"`          
    MotherName   string 	`gorm:"not null" json:"mother_name"`           
    Gender       string 	`gorm:"type:enum('L','P');not null" json:"gender"`
    Level      	 string 	`gorm:"type:enum('X','XI', 'XII');" json:"level"`              
}
