package repository

import (
	"github.com/Tretorhate/university-management-system/internal/domain"
)

type TeacherRepository struct {
	*Repository
}

func NewTeacherRepository(repo *Repository) *TeacherRepository {
	return &TeacherRepository{Repository: repo}
}

func (r *TeacherRepository) Create(teacher *domain.Teacher) error {
	return r.db.Create(teacher).Error
}

func (r *TeacherRepository) FindAll() ([]domain.Teacher, error) {
	var teachers []domain.Teacher
	if err := r.db.Preload("User").Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *TeacherRepository) FindByID(id uint) (*domain.Teacher, error) {
	var teacher domain.Teacher
	if err := r.db.Preload("User").First(&teacher, id).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *TeacherRepository) FindByEmployeeID(employeeID string) (*domain.Teacher, error) {
	var teacher domain.Teacher
	if err := r.db.Preload("User").Where("employee_id = ?", employeeID).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *TeacherRepository) FindByUserID(userID uint) (*domain.Teacher, error) {
	var teacher domain.Teacher
	if err := r.db.Preload("User").Where("user_id = ?", userID).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *TeacherRepository) Update(teacher *domain.Teacher) error {
	return r.db.Save(teacher).Error
}

func (r *TeacherRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Teacher{}, id).Error
}
