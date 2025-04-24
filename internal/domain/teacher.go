package domain

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      uint           `gorm:"not null" json:"userId"`
	User        User           `gorm:"foreignKey:UserID" json:"user"`
	EmployeeID  string         `gorm:"unique;not null" json:"employeeId"`
	Department  string         `gorm:"not null" json:"department"`
	Speciality  string         `gorm:"not null" json:"speciality"`
	JoiningDate time.Time      `gorm:"not null" json:"joiningDate"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
