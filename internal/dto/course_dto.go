package dto

import "time"

type CourseCreateDTO struct {
	Code        string    `json:"code" binding:"required,course_code"`
	Name        string    `json:"name" binding:"required,min=2,max=100"`
	Description string    `json:"description" binding:"max=500"`
	Credits     int       `json:"credits" binding:"required,min=1,max=6"`
	TeacherID   uint      `json:"teacherId" binding:"required"`
	StartDate   time.Time `json:"startDate" binding:"required"`
	EndDate     time.Time `json:"endDate" binding:"required,date_range"`
}

type CourseResponseDTO struct {
	ID          uint      `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Credits     int       `json:"credits"`
	TeacherID   uint      `json:"teacherId"`
	TeacherName string    `json:"teacherName"` // Combining teacher first and last name
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

type CourseUpdateDTO struct {
	Name        string    `json:"name" binding:"omitempty,min=2,max=100"`
	Description string    `json:"description" binding:"omitempty,max=500"`
	Credits     int       `json:"credits" binding:"omitempty,min=1,max=6"`
	TeacherID   uint      `json:"teacherId" binding:"omitempty"`
	StartDate   time.Time `json:"startDate" binding:"omitempty"`
	EndDate     time.Time `json:"endDate" binding:"omitempty,date_range"`
}
