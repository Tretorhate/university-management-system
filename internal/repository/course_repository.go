package repository

import (
	"github.com/Tretorhate/university-management-system/internal/domain"
)

type CourseRepository struct {
	*Repository
}

func NewCourseRepository(repo *Repository) *CourseRepository {
	return &CourseRepository{Repository: repo}
}

func (r *CourseRepository) Create(course *domain.Course) error {
	return r.db.Create(course).Error
}

func (r *CourseRepository) FindAll() ([]domain.Course, error) {
	var courses []domain.Course
	if err := r.db.Preload("Teacher.User").Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepository) FindByID(id uint) (*domain.Course, error) {
	var course domain.Course
	if err := r.db.Preload("Teacher.User").First(&course, id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepository) FindByCode(code string) (*domain.Course, error) {
	var course domain.Course
	if err := r.db.Preload("Teacher.User").Where("code = ?", code).First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepository) FindByTeacherID(teacherID uint) ([]domain.Course, error) {
	var courses []domain.Course
	if err := r.db.Preload("Teacher.User").Where("teacher_id = ?", teacherID).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepository) Update(course *domain.Course) error {
	return r.db.Save(course).Error
}

func (r *CourseRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Course{}, id).Error
}
