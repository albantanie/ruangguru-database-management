package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type LessonRepo struct {
	db *gorm.DB
}

func NewLessonRepo(db *gorm.DB) LessonRepo {
	return LessonRepo{db}
}

func (l LessonRepo) Init(data []model.Lesson) error {
	for _, lesson := range data {
		if err := l.db.Create(&lesson).Error; err != nil {
			return err
		}
	}
	return nil // TODO: replace this
}
