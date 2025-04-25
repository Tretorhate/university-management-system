package service

import (
	"errors"

	"github.com/Tretorhate/university-management-system/internal/domain"
	"github.com/Tretorhate/university-management-system/internal/dto"
	"github.com/Tretorhate/university-management-system/internal/repository"
	"github.com/Tretorhate/university-management-system/internal/service/factory"
	"github.com/Tretorhate/university-management-system/internal/service/strategy"
)

type CourseService struct {
	courseRepo           *repository.CourseRepository
	teacherRepo          *repository.TeacherRepository
	enrollRepo           *repository.EnrollmentRepository
	courseFactory        *factory.CourseFactory
	courseDTOFactory     *factory.CourseResponseDTOFactory
}

func NewCourseService(courseRepo *repository.CourseRepository, teacherRepo *repository.TeacherRepository) *CourseService {
	return &CourseService{
		courseRepo:       courseRepo,
		teacherRepo:      teacherRepo,
		courseFactory:    factory.NewCourseFactory(),
		courseDTOFactory: factory.NewCourseResponseDTOFactory(),
	}
}

func (s *CourseService) Create(req *dto.CourseCreateDTO) (*dto.CourseResponseDTO, error) {
	// Check if course code already exists
	existingCourse, _ := s.courseRepo.FindByCode(req.Code)
	if existingCourse != nil {
		return nil, errors.New("course with this code already exists")
	}

	// Verify teacher exists
	teacher, err := s.teacherRepo.FindByID(req.TeacherID)
	if err != nil {
		return nil, errors.New("teacher not found")
	}

	// Create course using factory
	course := s.courseFactory.CreateFromDTO(req)

	if err := s.courseRepo.Create(course); err != nil {
		return nil, err
	}

	// Set the Teacher field for the DTO conversion
	course.Teacher = *teacher
	
	return s.courseDTOFactory.CreateFromEntity(course), nil
}

func (s *CourseService) GetAll(sortBy string) ([]dto.CourseResponseDTO, error) {
	courses, err := s.courseRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// Apply sorting strategy
	var sortedCourses []domain.Course
	switch sortBy {
	case "date":
		strategy := &strategy.CourseByDateStrategy{}
		sortedCourses = strategy.Sort(courses)
	case "name":
		strategy := &strategy.CourseByNameStrategy{}
		sortedCourses = strategy.Sort(courses)
	case "students":
		// Get enrollment counts for each course
		enrollmentCounts := make(map[uint]int)
		for _, course := range courses {
			count, _ := s.enrollRepo.CountByCourseID(course.ID)
			enrollmentCounts[course.ID] = count
		}
		strategy := strategy.NewCourseByStudentCountStrategy(enrollmentCounts)
		sortedCourses = strategy.Sort(courses)
	default:
		sortedCourses = courses
	}

	var dtos []dto.CourseResponseDTO
	for _, course := range sortedCourses {
		dtos = append(dtos, *s.courseDTOFactory.CreateFromEntity(&course))
	}

	return dtos, nil
}

func (s *CourseService) GetByID(id uint) (*dto.CourseResponseDTO, error) {
	course, err := s.courseRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.courseDTOFactory.CreateFromEntity(course), nil
}

func (s *CourseService) Update(id uint, req *dto.CourseUpdateDTO) (*dto.CourseResponseDTO, error) {
	course, err := s.courseRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update course info
	if req.Name != "" {
		course.Name = req.Name
	}
	if req.Description != "" {
		course.Description = req.Description
	}
	if req.Credits != 0 {
		course.Credits = req.Credits
	}
	if req.TeacherID != 0 && req.TeacherID != course.TeacherID {
		// Verify new teacher exists
		teacher, err := s.teacherRepo.FindByID(req.TeacherID)
		if err != nil {
			return nil, errors.New("teacher not found")
		}
		course.TeacherID = req.TeacherID
		course.Teacher = *teacher
	}
	if !req.StartDate.IsZero() {
		course.StartDate = req.StartDate
	}
	if !req.EndDate.IsZero() {
		course.EndDate = req.EndDate
	}

	if err := s.courseRepo.Update(course); err != nil {
		return nil, err
	}

	return s.courseDTOFactory.CreateFromEntity(course), nil
}

func (s *CourseService) Delete(id uint) error {
	return s.courseRepo.Delete(id)
}