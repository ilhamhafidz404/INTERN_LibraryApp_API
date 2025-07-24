package models

import "gorm.io/gorm"

type Book struct {
    gorm.Model      
    Title       string 	`gorm:"type:varchar(100);unique;not null" json:"title"`          
    Cover       string 	`gorm:"type:varchar(255);not null" json:"cover"`          
    Publisher   string 	`gorm:"type:varchar(50);not null" json:"publisher"`               
    Author 		string 	`gorm:"type:varchar(50);not null" json:"author"`      
    ISBN  		string 	`gorm:"type:varchar(13);unique;not null" json:"isbn"`         
    Year 		uint16 	`gorm:"not null" json:"year"`          
    Total       int64 	`gorm:"not null" json:"total"`          
    CreatedBy   uint    `gorm:"not null" json:"created_by"`
    Admin       Admin   `gorm:"foreignKey:CreatedBy" json:"admin"`   
}
