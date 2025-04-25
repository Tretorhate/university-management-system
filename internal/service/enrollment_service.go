package service

import (
	"github.com/Tretorhate/university-management-system/internal/domain"
	"github.com/Tretorhate/university-management-system/internal/dto"
	"github.com/Tretorhate/university-management-system/internal/repository"
	"github.com/Tretorhate/university-management-system/internal/service/factory"
	"github.com/Tretorhate/university-management-system/pkg/errors"
)

type EnrollmentService struct {
	enrollmentRepo           *repository.EnrollmentRepository
	studentRepo              *repository.StudentRepository
	courseRepo               *repository.CourseRepository
	enrollmentFactory        *factory.EnrollmentFactory
	enrollmentDTOFactory     *factory.EnrollmentResponseDTOFactory
}

func NewEnrollmentService(enrollmentRepo *repository.EnrollmentRepository, studentRepo *repository.StudentRepository, courseRepo *repository.CourseRepository) *EnrollmentService {
	return &EnrollmentService{
		enrollmentRepo:       enrollmentRepo,
		studentRepo:          studentRepo,
		courseRepo:           courseRepo,
		enrollmentFactory:    factory.NewEnrollmentFactory(),
		enrollmentDTOFactory: factory.NewEnrollmentResponseDTOFactory(),
	}
}

func (s *EnrollmentService) Create(req *dto.EnrollmentCreateDTO) (*dto.EnrollmentResponseDTO, error) {
	// Verify student exists
	student, err := s.studentRepo.FindByID(req.StudentID)
	if err != nil {
		return nil, errors.NotFound("Student not found", err)
	}

	// Verify course exists
	course, err := s.courseRepo.FindByID(req.CourseID)
	if err != nil {
		return nil, errors.NotFound("Course not found", err)
	}

	// Check if enrollment already exists
	enrollments, _ := s.enrollmentRepo.FindByStudentID(req.StudentID)
	for _, e := range enrollments {
		if e.CourseID == req.CourseID {
			return nil, errors.BadRequest("Student is already enrolled in this course", nil)
		}
	}

	// Create enrollment using factory
	enrollment := s.enrollmentFactory.CreateFromDTO(req)

	if err := s.enrollmentRepo.Create(enrollment); err != nil {
		return nil, errors.InternalServerError("Failed to create enrollment", err)
	}

	// Set the Student and Course fields for the DTO conversion
	enrollment.Student = *student
	enrollment.Course = *course
	
	return s.enrollmentDTOFactory.CreateFromEntity(enrollment), nil
}

func (s *EnrollmentService) GetAll() ([]dto.EnrollmentResponseDTO, error) {
	enrollments, err := s.enrollmentRepo.FindAll()
	if err != nil {
		return nil, errors.InternalServerError("Failed to retrieve enrollments", err)
	}

	var dtos []dto.EnrollmentResponseDTO
	for _, enrollment := range enrollments {
		dtos = append(dtos, *s.enrollmentDTOFactory.CreateFromEntity(&enrollment))
	}

	return dtos, nil
}

func (s *EnrollmentService) GetByID(id uint) (*dto.EnrollmentResponseDTO, error) {
	enrollment, err := s.enrollmentRepo.FindByID(id)
	if err != nil {
		return nil, errors.NotFound("Enrollment not found", err)
	}

	return s.enrollmentDTOFactory.CreateFromEntity(enrollment), nil
}

func (s *EnrollmentService) Update(id uint, req *dto.EnrollmentUpdateDTO) (*dto.EnrollmentResponseDTO, error) {
	enrollment, err := s.enrollmentRepo.FindByID(id)
	if err != nil {
		return nil, errors.NotFound("Enrollment not found", err)
	}

	// Update enrollment info
	if req.Grade != nil {
		enrollment.Grade = req.Grade
	}
	if !req.EnrollDate.IsZero() {
		enrollment.EnrollDate = req.EnrollDate
	}

	if err := s.enrollmentRepo.Update(enrollment); err != nil {
		return nil, errors.InternalServerError("Failed to update enrollment", err)
	}

	return s.enrollmentDTOFactory.CreateFromEntity(enrollment), nil
}

func (s *EnrollmentService) Delete(id uint) error {
	// Check if enrollment exists
	_, err := s.enrollmentRepo.FindByID(id)
	if err != nil {
		return errors.NotFound("Enrollment not found", err)
	}

	if err := s.enrollmentRepo.Delete(id); err != nil {
		return errors.InternalServerError("Failed to delete enrollment", err)
	}

	return nil
}

func (s *EnrollmentService) GetByStudentID(studentID uint) ([]dto.EnrollmentResponseDTO, error) {
	// Verify student exists
	_, err := s.studentRepo.FindByID(studentID)
	if err != nil {
		return nil, errors.NotFound("Student not found", err)
	}

	enrollments, err := s.enrollmentRepo.FindByStudentID(studentID)
	if err != nil {
		return nil, errors.InternalServerError("Failed to retrieve enrollments", err)
	}

	var dtos []dto.EnrollmentResponseDTO
	for _, enrollment := range enrollments {
		dtos = append(dtos, *s.enrollmentDTOFactory.CreateFromEntity(&enrollment))
	}

	return dtos, nil
}

func (s *EnrollmentService) GetByCourseID(courseID uint) ([]dto.EnrollmentResponseDTO, error) {
	// Verify course exists
	_, err := s.courseRepo.FindByID(courseID)
	if err != nil {
		return nil, errors.NotFound("Course not found", err)
	}

	enrollments, err := s.enrollmentRepo.FindByCourseID(courseID)
	if err != nil {
		return nil, errors.InternalServerError("Failed to retrieve enrollments", err)
	}

	var dtos []dto.EnrollmentResponseDTO
	for _, enrollment := range enrollments {
		dtos = append(dtos, *s.enrollmentDTOFactory.CreateFromEntity(&enrollment))
	}

	return dtos, nil
}

func (s *EnrollmentService) GetByStudentAndCourseID(studentID, courseID uint) ([]dto.EnrollmentResponseDTO, error) {
	// Verify student exists
	_, err := s.studentRepo.FindByID(studentID)
	if err != nil {
		return nil, errors.NotFound("Student not found", err)
	}

	// Verify course exists
	_, err = s.courseRepo.FindByID(courseID)
	if err != nil {
		return nil, errors.NotFound("Course not found", err)
	}

	// First get all enrollments by studentID
	enrollments, err := s.enrollmentRepo.FindByStudentID(studentID)
	if err != nil {
		return nil, errors.InternalServerError("Failed to retrieve enrollments", err)
	}

	// Filter by courseID
	var filtered []domain.Enrollment
	for _, enrollment := range enrollments {
		if enrollment.CourseID == courseID {
			filtered = append(filtered, enrollment)
		}
	}

	var dtos []dto.EnrollmentResponseDTO
	for _, enrollment := range filtered {
		dtos = append(dtos, *s.enrollmentDTOFactory.CreateFromEntity(&enrollment))
	}

	return dtos, nil
}