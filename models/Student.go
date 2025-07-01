package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
    gorm.Model
    NISN         string 	`gorm:"type:varchar(10);unique;not null" json:"nisn"`          
    NIK          string 	`gorm:"type:varchar(16);unique" json:"nik"`          
    Name         string 	`gorm:"type:varchar(50);not null" json:"name"`               
    Password     string 	`gorm:"type:varchar(255);not null" json:"password"`               
    PlaceOfBirth string 	`gorm:"type:varchar(30);not null" json:"place_of_birth"`      
    DateOfBirth  time.Time 	`gorm:"type:datetime;not null" json:"date_of_birth"`          
    MotherName   string 	`gorm:"type:varchar(30)" json:"mother_name"`           
    Gender       string 	`gorm:"type:enum('M','F');not null" json:"gender"`
    Level      	 string 	`gorm:"type:enum('X','XI', 'XII');not null" json:"level"`              
}
