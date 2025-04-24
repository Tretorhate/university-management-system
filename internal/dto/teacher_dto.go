package dto

import "time"

type TeacherCreateDTO struct {
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required,password"`
	FirstName   string    `json:"firstName" binding:"required,min=2,max=50"`
	LastName    string    `json:"lastName" binding:"required,min=2,max=50"`
	EmployeeID  string    `json:"employeeId" binding:"required,employee_id"`
	Department  string    `json:"department" binding:"required,min=2,max=100"`
	Speciality  string    `json:"speciality" binding:"required,min=2,max=100"`
	JoiningDate time.Time `json:"joiningDate" binding:"required"`
}

type TeacherResponseDTO struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"userId"`
	Email       string    `json:"email"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	EmployeeID  string    `json:"employeeId"`
	Department  string    `json:"department"`
	Speciality  string    `json:"speciality"`
	JoiningDate time.Time `json:"joiningDate"`
}

type TeacherUpdateDTO struct {
	FirstName   string    `json:"firstName" binding:"omitempty,min=2,max=50"`
	LastName    string    `json:"lastName" binding:"omitempty,min=2,max=50"`
	Department  string    `json:"department" binding:"omitempty,min=2,max=100"`
	Speciality  string    `json:"speciality" binding:"omitempty,min=2,max=100"`
	JoiningDate time.Time `json:"joiningDate" binding:"omitempty"`
}
