package domain

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Code        string         `gorm:"unique;not null" json:"code"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Credits     int            `gorm:"not null" json:"credits"`
	TeacherID   uint           `gorm:"not null" json:"teacherId"`
	Teacher     Teacher        `gorm:"foreignKey:TeacherID" json:"teacher"`
	StartDate   time.Time      `gorm:"not null" json:"startDate"`
	EndDate     time.Time      `gorm:"not null" json:"endDate"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
