package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	return t.db.Create(&data).Error
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	var teachers []model.Teacher
	if err := t.db.Unscoped().Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (t TeacherRepo) Update(id uint, name string) error {
	return t.db.Model(&model.Teacher{}).Where("id = ?", id).Update("name", name).Error
}

func (t TeacherRepo) Delete(id uint) error {
	return t.db.Delete(&model.Teacher{}, id).Error
}
