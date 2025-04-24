package factory

import (
	"github.com/Tretorhate/university-management-system/internal/domain"
	"github.com/Tretorhate/university-management-system/internal/dto"
)

// UserDTOFactory is a factory for creating UserDTO objects
type UserDTOFactory struct{}

func NewUserDTOFactory() *UserDTOFactory {
	return &UserDTOFactory{}
}

func (f *UserDTOFactory) CreateFromEntity(user *domain.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      string(user.Role),
	}
}

// StudentDTOFactory is a factory for creating StudentResponseDTO objects
type StudentDTOFactory struct{}

func NewStudentDTOFactory() *StudentDTOFactory {
	return &StudentDTOFactory{}
}

func (f *StudentDTOFactory) CreateFromEntity(student *domain.Student) *dto.StudentResponseDTO {
	return &dto.StudentResponseDTO{
		ID:         student.ID,
		UserID:     student.UserID,
		Email:      student.User.Email,
		FirstName:  student.User.FirstName,
		LastName:   student.User.LastName,
		StudentID:  student.StudentID,
		EnrollYear: student.EnrollYear,
		Major:      student.Major,
	}
}

// CourseFactory is a factory for creating Course entities
type CourseFactory struct{}

func NewCourseFactory() *CourseFactory {
	return &CourseFactory{}
}

func (f *CourseFactory) CreateFromDTO(dto *dto.CourseCreateDTO, teacherID uint) *domain.Course {
	return &domain.Course{
		Code:        dto.Code,
		Name:        dto.Name,
		Description: dto.Description,
		Credits:     dto.Credits,
		TeacherID:   teacherID,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}
}
