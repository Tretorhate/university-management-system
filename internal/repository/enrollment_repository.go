package repository

import (
	"github.com/Tretorhate/university-management-system/internal/domain"
)

type EnrollmentRepository struct {
	*Repository
}

func NewEnrollmentRepository(repo *Repository) *EnrollmentRepository {
	return &EnrollmentRepository{Repository: repo}
}

func (r *EnrollmentRepository) Create(enrollment *domain.Enrollment) error {
	return r.db.Create(enrollment).Error
}

func (r *EnrollmentRepository) FindAll() ([]domain.Enrollment, error) {
	var enrollments []domain.Enrollment
	if err := r.db.Preload("Student.User").Preload("Course.Teacher.User").Find(&enrollments).Error; err != nil {
		return nil, err
	}
	return enrollments, nil
}

func (r *EnrollmentRepository) FindByID(id uint) (*domain.Enrollment, error) {
	var enrollment domain.Enrollment
	if err := r.db.Preload("Student.User").Preload("Course.Teacher.User").First(&enrollment, id).Error; err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (r *EnrollmentRepository) FindByStudentID(studentID uint) ([]domain.Enrollment, error) {
	var enrollments []domain.Enrollment
	if err := r.db.Preload("Student.User").Preload("Course.Teacher.User").Where("student_id = ?", studentID).Find(&enrollments).Error; err != nil {
		return nil, err
	}
	return enrollments, nil
}

func (r *EnrollmentRepository) FindByCourseID(courseID uint) ([]domain.Enrollment, error) {
	var enrollments []domain.Enrollment
	if err := r.db.Preload("Student.User").Preload("Course.Teacher.User").Where("course_id = ?", courseID).Find(&enrollments).Error; err != nil {
		return nil, err
	}
	return enrollments, nil
}

func (r *EnrollmentRepository) CountByCourseID(courseID uint) (int, error) {
	var count int64
	if err := r.db.Model(&domain.Enrollment{}).Where("course_id = ?", courseID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *EnrollmentRepository) Update(enrollment *domain.Enrollment) error {
	return r.db.Save(enrollment).Error
}

func (r *EnrollmentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Enrollment{}, id).Error
}

func (r *EnrollmentRepository) DeleteByStudentAndCourse(studentID, courseID uint) error {
	return r.db.Where("student_id = ? AND course_id = ?", studentID, courseID).Delete(&domain.Enrollment{}).Error
}
