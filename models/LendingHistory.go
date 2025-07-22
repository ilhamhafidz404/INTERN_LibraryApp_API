package models

import (
	"time"

	"gorm.io/gorm"
)

type LendingHistory struct {
    gorm.Model
    StartDate  time.Time
    EndDate    time.Time
    BookID     uint    `gorm:"not null"`
    Book       Book    `gorm:"foreignKey:BookID"`
    StudentID  uint    `gorm:"not null"`
    Student    Student `gorm:"foreignKey:StudentID"`
    Status     string  `gorm:"type:enum('loaned','returned','late returned');default:'loaned';not null"`
}

