package service

import (
	"errors"

	"github.com/Tretorhate/university-management-system/internal/domain"
	"github.com/Tretorhate/university-management-system/internal/dto"
	"github.com/Tretorhate/university-management-system/internal/repository"
	"github.com/Tretorhate/university-management-system/internal/service/factory"
	"golang.org/x/crypto/bcrypt"
)

type TeacherService struct {
	teacherRepo      *repository.TeacherRepository
	userRepo         *repository.UserRepository
	teacherFactory   *factory.TeacherFactory
	teacherDTOFactory *factory.TeacherDTOFactory
}

func NewTeacherService(teacherRepo *repository.TeacherRepository, userRepo *repository.UserRepository) *TeacherService {
	return &TeacherService{
		teacherRepo:      teacherRepo,
		userRepo:         userRepo,
		teacherFactory:   factory.NewTeacherFactory(),
		teacherDTOFactory: factory.NewTeacherDTOFactory(),
	}
}

func (s *TeacherService) Create(req *dto.TeacherCreateDTO) (*dto.TeacherResponseDTO, error) {
	// Check if email already exists
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Check if employee ID already exists
	existingTeacher, _ := s.teacherRepo.FindByEmployeeID(req.EmployeeID)
	if existingTeacher != nil {
		return nil, errors.New("teacher with this employee ID already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user with TEACHER role
	user := &domain.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      domain.RoleTeacher,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Create teacher using factory
	teacher := s.teacherFactory.CreateFromDTO(req, user.ID)

	if err := s.teacherRepo.Create(teacher); err != nil {
		// Rollback user creation if teacher creation fails
		_ = s.userRepo.Delete(user.ID)
		return nil, err
	}

	// Setting the User field for the DTO conversion
	teacher.User = *user
	
	// Using factory to create response DTO
	return s.teacherDTOFactory.CreateFromEntity(teacher), nil
}

func (s *TeacherService) GetAll() ([]dto.TeacherResponseDTO, error) {
	teachers, err := s.teacherRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var dtos []dto.TeacherResponseDTO
	for _, teacher := range teachers {
		dtos = append(dtos, *s.teacherDTOFactory.CreateFromEntity(&teacher))
	}

	return dtos, nil
}

func (s *TeacherService) GetByID(id uint) (*dto.TeacherResponseDTO, error) {
	teacher, err := s.teacherRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.teacherDTOFactory.CreateFromEntity(teacher), nil
}

func (s *TeacherService) Update(id uint, req *dto.TeacherUpdateDTO) (*dto.TeacherResponseDTO, error) {
	teacher, err := s.teacherRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(teacher.UserID)
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

	// Update teacher info
	if req.Department != "" {
		teacher.Department = req.Department
	}
	if req.Speciality != "" {
		teacher.Speciality = req.Speciality
	}
	if !req.JoiningDate.IsZero() {
		teacher.JoiningDate = req.JoiningDate
	}
	if err := s.teacherRepo.Update(teacher); err != nil {
		return nil, err
	}

	// Update the User field for the DTO conversion
	teacher.User = *user
	
	return s.teacherDTOFactory.CreateFromEntity(teacher), nil
}

func (s *TeacherService) Delete(id uint) error {
	teacher, err := s.teacherRepo.FindByID(id)
	if err != nil {
		return err
	}

	// First delete teacher
	if err := s.teacherRepo.Delete(id); err != nil {
		return err
	}

	// Then delete user
	return s.userRepo.Delete(teacher.UserID)
}