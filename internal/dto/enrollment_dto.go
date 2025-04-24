// internal/dto/enrollment_dto.go
package dto

import "time"

type EnrollmentCreateDTO struct {
	StudentID  uint      `json:"studentId" binding:"required"`
	CourseID   uint      `json:"courseId" binding:"required"`
	EnrollDate time.Time `json:"enrollDate" binding:"required"`
}

type EnrollmentResponseDTO struct {
	ID          uint      `json:"id"`
	StudentID   uint      `json:"studentId"`
	StudentName string    `json:"studentName"` // Combined first and last name
	CourseID    uint      `json:"courseId"`
	CourseName  string    `json:"courseName"`
	CourseCode  string    `json:"courseCode"`
	Grade       *float64  `json:"grade"`
	EnrollDate  time.Time `json:"enrollDate"`
}

type EnrollmentUpdateDTO struct {
	Grade      *float64  `json:"grade" binding:"omitempty,min=0,max=100"`
	EnrollDate time.Time `json:"enrollDate" binding:"omitempty"`
}
