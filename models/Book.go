package models

import "gorm.io/gorm"

type Book struct {
    gorm.Model
    Code        string 	`gorm:"type:varchar(7);unique;not null" json:"code"`          
    Title       string 	`gorm:"type:varchar(100);unique;not null" json:"title"`          
    Publisher   string 	`gorm:"type:varchar(50);not null" json:"publisher"`               
    Author 		string 	`gorm:"type:varchar(50);not null" json:"author"`      
    ISBN  		string 	`gorm:"type:varchar(13);unique;not null" json:"isbn"`         
    Year 		uint16 	`gorm:"not null" json:"year"`          
    Total       int64 	`gorm:"not null" json:"total"`          
}
