package models

import (
	"gorm.io/gorm"
)

type Admin struct {
    gorm.Model             
    Name         string 	`gorm:"type:varchar(50);not null"`               
    Username     string 	`gorm:"type:varchar(50);not null"`               
    Password     string 	`gorm:"type:varchar(255);not null"`     
    Level      	 string 	`gorm:"type:enum('super_admin','admin');not null;default:admin"`    
    Books        []Book     `gorm:"foreignKey:CreatedBy"`          
}
