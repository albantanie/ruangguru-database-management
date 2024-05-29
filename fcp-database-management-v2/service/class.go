package service

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
)

type ClassService interface {
	FetchAll() ([]model.Class, error)
}

type classService struct {
	classRepository repository.ClassRepository
}

func NewClassService(classRepository repository.ClassRepository) ClassService {
	return &classService{classRepository}
}

func (s *classService) FetchAll() ([]model.Class, error) {
	classes, err := s.classRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return classes, nil
}
