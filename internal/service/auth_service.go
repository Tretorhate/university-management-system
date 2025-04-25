package service

import (
	"github.com/Tretorhate/university-management-system/internal/domain"
	"github.com/Tretorhate/university-management-system/internal/dto"
	"github.com/Tretorhate/university-management-system/internal/repository"
	"github.com/Tretorhate/university-management-system/internal/service/factory"
	"github.com/Tretorhate/university-management-system/pkg/auth"
	"github.com/Tretorhate/university-management-system/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo     *repository.UserRepository
	jwtService   *auth.JWTService
	userDTOFactory *factory.UserDTOFactory
}

func NewAuthService(userRepo *repository.UserRepository, jwtService *auth.JWTService) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		jwtService:   jwtService,
		userDTOFactory: factory.NewUserDTOFactory(),
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.BadRequest("User with this email already exists", nil)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.InternalServerError("Failed to hash password", err)
	}

	// Create user
	user := &domain.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      domain.Role(req.Role),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.InternalServerError("Failed to create user", err)
	}

	// Generate token
	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, errors.InternalServerError("Failed to generate token", err)
	}

	// Use factory to create user DTO
	userDTO := s.userDTOFactory.CreateFromEntity(user)

	return &dto.AuthResponse{
		Token: token,
		User:  *userDTO,
	}, nil
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	// Find user
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.Unauthorized("Invalid email or password", nil)
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.Unauthorized("Invalid email or password", nil)
	}

	// Generate token
	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, errors.InternalServerError("Failed to generate token", err)
	}

	// Use factory to create user DTO
	userDTO := s.userDTOFactory.CreateFromEntity(user)

	return &dto.AuthResponse{
		Token: token,
		User:  *userDTO,
	}, nil
}