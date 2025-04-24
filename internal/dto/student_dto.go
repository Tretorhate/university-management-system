package dto

type StudentCreateDTO struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,password"`
	FirstName  string `json:"firstName" binding:"required,min=2,max=50"`
	LastName   string `json:"lastName" binding:"required,min=2,max=50"`
	StudentID  string `json:"studentId" binding:"required,student_id"`
	EnrollYear int    `json:"enrollYear" binding:"required,min=2000,max=2100"`
	Major      string `json:"major" binding:"required,min=2,max=100"`
}

type StudentResponseDTO struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"userId"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	StudentID  string `json:"studentId"`
	EnrollYear int    `json:"enrollYear"`
	Major      string `json:"major"`
}

type StudentUpdateDTO struct {
	FirstName  string `json:"firstName" binding:"omitempty,min=2,max=50"`
	LastName   string `json:"lastName" binding:"omitempty,min=2,max=50"`
	EnrollYear int    `json:"enrollYear" binding:"omitempty,min=2000,max=2100"`
	Major      string `json:"major" binding:"omitempty,min=2,max=100"`
}
