package domain

import (
	"time"

	"gorm.io/gorm"
)

type Enrollment struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	StudentID  uint           `gorm:"not null" json:"studentId"`
	Student    Student        `gorm:"foreignKey:StudentID;references:ID" json:"student"`
	CourseID   uint           `gorm:"not null" json:"courseId"`
	Course     Course         `gorm:"foreignKey:CourseID;references:ID" json:"course"`
	Grade      *float64       `gorm:"type:double precision" json:"grade"`
	EnrollDate time.Time      `gorm:"not null" json:"enrollDate"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
