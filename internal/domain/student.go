package domain

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	UserID     uint           `gorm:"not null" json:"userId"`
	User       User           `gorm:"foreignKey:UserID" json:"user"`
	StudentID  string         `gorm:"unique;not null" json:"studentId"`
	EnrollYear int            `gorm:"not null" json:"enrollYear"`
	Major      string         `gorm:"not null" json:"major"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
