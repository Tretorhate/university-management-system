package strategy

import (
	"sort"

	"github.com/Tretorhate/university-management-system/internal/domain"
)

// CourseSortStrategy defines the interface for sorting courses
type CourseSortStrategy interface {
	Sort(courses []domain.Course) []domain.Course
}

// CourseByDateStrategy sorts courses by start date
type CourseByDateStrategy struct{}

func (s *CourseByDateStrategy) Sort(courses []domain.Course) []domain.Course {
	sorted := make([]domain.Course, len(courses))
	copy(sorted, courses)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].StartDate.Before(sorted[j].StartDate)
	})

	return sorted
}


type CourseByStudentCountStrategy struct {
	EnrollmentCounts map[uint]int 
}

func NewCourseByStudentCountStrategy(enrollmentCounts map[uint]int) *CourseByStudentCountStrategy {
	return &CourseByStudentCountStrategy{
		EnrollmentCounts: enrollmentCounts,
	}
}

func (s *CourseByStudentCountStrategy) Sort(courses []domain.Course) []domain.Course {
	sorted := make([]domain.Course, len(courses))
	copy(sorted, courses)

	sort.Slice(sorted, func(i, j int) bool {
		countI := s.EnrollmentCounts[sorted[i].ID]
		countJ := s.EnrollmentCounts[sorted[j].ID]
		return countI > countJ // Descending order
	})

	return sorted
}

// CourseByNameStrategy sorts courses alphabetically by name
type CourseByNameStrategy struct{}

func (s *CourseByNameStrategy) Sort(courses []domain.Course) []domain.Course {
	sorted := make([]domain.Course, len(courses))
	copy(sorted, courses)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Name < sorted[j].Name
	})

	return sorted
}

// CourseSorter implements the strategy pattern
type CourseSorter struct {
	strategy CourseSortStrategy
}

func NewCourseSorter(strategy CourseSortStrategy) *CourseSorter {
	return &CourseSorter{
		strategy: strategy,
	}
}

func (s *CourseSorter) SetStrategy(strategy CourseSortStrategy) {
	s.strategy = strategy
}

func (s *CourseSorter) Sort(courses []domain.Course) []domain.Course {
	return s.strategy.Sort(courses)
}
