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

// TeacherDTOFactory is a factory for creating TeacherResponseDTO objects
type TeacherDTOFactory struct{}

func NewTeacherDTOFactory() *TeacherDTOFactory {
	return &TeacherDTOFactory{}
}

func (f *TeacherDTOFactory) CreateFromEntity(teacher *domain.Teacher) *dto.TeacherResponseDTO {
	return &dto.TeacherResponseDTO{
		ID:          teacher.ID,
		UserID:      teacher.UserID,
		Email:       teacher.User.Email,
		FirstName:   teacher.User.FirstName,
		LastName:    teacher.User.LastName,
		EmployeeID:  teacher.EmployeeID,
		Department:  teacher.Department,
		Speciality:  teacher.Speciality,
		JoiningDate: teacher.JoiningDate,
	}
}

// CourseResponseDTOFactory is a factory for creating CourseResponseDTO objects
type CourseResponseDTOFactory struct{}

func NewCourseResponseDTOFactory() *CourseResponseDTOFactory {
	return &CourseResponseDTOFactory{}
}

func (f *CourseResponseDTOFactory) CreateFromEntity(course *domain.Course) *dto.CourseResponseDTO {
	teacherName := ""
	if course.Teacher.User.FirstName != "" {
		teacherName = course.Teacher.User.FirstName + " " + course.Teacher.User.LastName
	}
	
	return &dto.CourseResponseDTO{
		ID:          course.ID,
		Code:        course.Code,
		Name:        course.Name,
		Description: course.Description,
		Credits:     course.Credits,
		TeacherID:   course.TeacherID,
		TeacherName: teacherName,
		StartDate:   course.StartDate,
		EndDate:     course.EndDate,
	}
}

// EnrollmentResponseDTOFactory is a factory for creating EnrollmentResponseDTO objects
type EnrollmentResponseDTOFactory struct{}

func NewEnrollmentResponseDTOFactory() *EnrollmentResponseDTOFactory {
	return &EnrollmentResponseDTOFactory{}
}

func (f *EnrollmentResponseDTOFactory) CreateFromEntity(enrollment *domain.Enrollment) *dto.EnrollmentResponseDTO {
	studentName := enrollment.Student.User.FirstName + " " + enrollment.Student.User.LastName
	
	return &dto.EnrollmentResponseDTO{
		ID:          enrollment.ID,
		StudentID:   enrollment.StudentID,
		StudentName: studentName,
		CourseID:    enrollment.CourseID,
		CourseName:  enrollment.Course.Name,
		CourseCode:  enrollment.Course.Code,
		Grade:       enrollment.Grade,
		EnrollDate:  enrollment.EnrollDate,
	}
}

// CourseFactory is a factory for creating Course entities from DTOs
type CourseFactory struct{}

func NewCourseFactory() *CourseFactory {
	return &CourseFactory{}
}

func (f *CourseFactory) CreateFromDTO(dto *dto.CourseCreateDTO) *domain.Course {
	return &domain.Course{
		Code:        dto.Code,
		Name:        dto.Name,
		Description: dto.Description,
		Credits:     dto.Credits,
		TeacherID:   dto.TeacherID,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}
}

// StudentFactory is a factory for creating Student entities from DTOs
type StudentFactory struct{}

func NewStudentFactory() *StudentFactory {
	return &StudentFactory{}
}

func (f *StudentFactory) CreateFromDTO(dto *dto.StudentCreateDTO, userID uint) *domain.Student {
	return &domain.Student{
		UserID:     userID,
		StudentID:  dto.StudentID,
		EnrollYear: dto.EnrollYear,
		Major:      dto.Major,
	}
}

// TeacherFactory is a factory for creating Teacher entities from DTOs
type TeacherFactory struct{}

func NewTeacherFactory() *TeacherFactory {
	return &TeacherFactory{}
}

func (f *TeacherFactory) CreateFromDTO(dto *dto.TeacherCreateDTO, userID uint) *domain.Teacher {
	return &domain.Teacher{
		UserID:      userID,
		EmployeeID:  dto.EmployeeID,
		Department:  dto.Department,
		Speciality:  dto.Speciality,
		JoiningDate: dto.JoiningDate,
	}
}

// EnrollmentFactory is a factory for creating Enrollment entities from DTOs
type EnrollmentFactory struct{}

func NewEnrollmentFactory() *EnrollmentFactory {
	return &EnrollmentFactory{}
}

func (f *EnrollmentFactory) CreateFromDTO(dto *dto.EnrollmentCreateDTO) *domain.Enrollment {
	return &domain.Enrollment{
		StudentID:  dto.StudentID,
		CourseID:   dto.CourseID,
		EnrollDate: dto.EnrollDate,
	}
}