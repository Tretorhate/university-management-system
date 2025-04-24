package service

import (
	"errors"

	"github.com/Tretorhate/university-management-system/internal/domain"
	"github.com/Tretorhate/university-management-system/internal/dto"
	"github.com/Tretorhate/university-management-system/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type StudentService struct {
	studentRepo *repository.StudentRepository
	userRepo    *repository.UserRepository
}

func NewStudentService(studentRepo *repository.StudentRepository, userRepo *repository.UserRepository) *StudentService {
	return &StudentService{
		studentRepo: studentRepo,
		userRepo:    userRepo,
	}
}

func (s *StudentService) Create(req *dto.StudentCreateDTO) (*dto.StudentResponseDTO, error) {
	// Check if email already exists
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Check if student ID already exists
	existingStudent, _ := s.studentRepo.FindByStudentID(req.StudentID)
	if existingStudent != nil {
		return nil, errors.New("student with this ID already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user with STUDENT role
	user := &domain.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      domain.RoleStudent,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Create student
	student := &domain.Student{
		UserID:     user.ID,
		StudentID:  req.StudentID,
		EnrollYear: req.EnrollYear,
		Major:      req.Major,
	}

	if err := s.studentRepo.Create(student); err != nil {
		// Rollback user creation if student creation fails
		_ = s.userRepo.Delete(user.ID)
		return nil, err
	}

	return &dto.StudentResponseDTO{
		ID:         student.ID,
		UserID:     user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		StudentID:  student.StudentID,
		EnrollYear: student.EnrollYear,
		Major:      student.Major,
	}, nil
}

func (s *StudentService) GetAll() ([]dto.StudentResponseDTO, error) {
	students, err := s.studentRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var dtos []dto.StudentResponseDTO
	for _, student := range students {
		dtos = append(dtos, dto.StudentResponseDTO{
			ID:         student.ID,
			UserID:     student.UserID,
			Email:      student.User.Email,
			FirstName:  student.User.FirstName,
			LastName:   student.User.LastName,
			StudentID:  student.StudentID,
			EnrollYear: student.EnrollYear,
			Major:      student.Major,
		})
	}

	return dtos, nil
}

func (s *StudentService) GetByID(id uint) (*dto.StudentResponseDTO, error) {
	student, err := s.studentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.StudentResponseDTO{
		ID:         student.ID,
		UserID:     student.UserID,
		Email:      student.User.Email,
		FirstName:  student.User.FirstName,
		LastName:   student.User.LastName,
		StudentID:  student.StudentID,
		EnrollYear: student.EnrollYear,
		Major:      student.Major,
	}, nil
}

func (s *StudentService) Update(id uint, req *dto.StudentUpdateDTO) (*dto.StudentResponseDTO, error) {
	student, err := s.studentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(student.UserID)
	if err != nil {
		return nil, err
	}

	// Update user info
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	// Update student info
	if req.EnrollYear != 0 {
		student.EnrollYear = req.EnrollYear
	}
	if req.Major != "" {
		student.Major = req.Major
	}
	if err := s.studentRepo.Update(student); err != nil {
		return nil, err
	}

	return &dto.StudentResponseDTO{
		ID:         student.ID,
		UserID:     user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		StudentID:  student.StudentID,
		EnrollYear: student.EnrollYear,
		Major:      student.Major,
	}, nil
}

func (s *StudentService) Delete(id uint) error {
	student, err := s.studentRepo.FindByID(id)
	if err != nil {
		return err
	}

	// First delete student
	if err := s.studentRepo.Delete(id); err != nil {
		return err
	}

	// Then delete user
	return s.userRepo.Delete(student.UserID)
}
