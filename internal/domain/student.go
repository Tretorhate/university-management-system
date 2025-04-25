package domain

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null" json:"userId"`
	User        User           `gorm:"foreignKey:UserID" json:"user"`
	StudentID   string         `gorm:"type:varchar(50);unique;not null" json:"studentId"`
	EnrollYear  int            `gorm:"not null" json:"enrollYear"`
	Major       string         `gorm:"type:varchar(255);not null" json:"major"`
	Enrollments []Enrollment   `gorm:"foreignKey:StudentID;references:ID" json:"enrollments"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
