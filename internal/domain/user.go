package domain

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin   Role = "ADMIN"
	RoleTeacher Role = "TEACHER"
	RoleStudent Role = "STUDENT"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	FirstName string         `gorm:"not null" json:"firstName"`
	LastName  string         `gorm:"not null" json:"lastName"`
	Role      Role           `gorm:"not null" json:"role"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
