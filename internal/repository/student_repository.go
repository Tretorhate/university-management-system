package repository

import (
	"github.com/Tretorhate/university-management-system/internal/domain"
)

type StudentRepository struct {
	*Repository
}

func NewStudentRepository(repo *Repository) *StudentRepository {
	return &StudentRepository{Repository: repo}
}

func (r *StudentRepository) Create(student *domain.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) FindAll() ([]domain.Student, error) {
	var students []domain.Student
	if err := r.db.Preload("User").Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *StudentRepository) FindByID(id uint) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.Preload("User").First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) FindByStudentID(studentID string) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.Preload("User").Where("student_id = ?", studentID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) FindByUserID(userID uint) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.Preload("User").Where("user_id = ?", userID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) Update(student *domain.Student) error {
	return r.db.Save(student).Error
}

func (r *StudentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Student{}, id).Error
}
