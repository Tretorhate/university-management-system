package domain

import (
	"time"

	"gorm.io/gorm"
)

type Enrollment struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	StudentID  uint           `gorm:"not null" json:"studentId"`
	Student    Student        `gorm:"foreignKey:StudentID" json:"student"`
	CourseID   uint           `gorm:"not null" json:"courseId"`
	Course     Course         `gorm:"foreignKey:CourseID" json:"course"`
	Grade      *float64       `json:"grade"`
	EnrollDate time.Time      `gorm:"not null" json:"enrollDate"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
