package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
	if err := s.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	if err := s.db.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	result := s.db.Create(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	result := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	result := s.db.Delete(&model.Student{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	var students []model.StudentClass
	if err := s.db.Table("students").
		Select("students.*, classes.name as class_name, classes.professor, classes.room_number").
		Joins("JOIN classes on students.class_id = classes.id").
		Scan(&students).Error; err != nil {
		return nil, err
	}
	if students == nil {
		students = []model.StudentClass{}
	}
	return &students, nil
}
