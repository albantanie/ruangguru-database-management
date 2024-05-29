package service

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
)

type StudentService interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
}

type studentService struct {
	studentRepository repository.StudentRepository
}

func NewStudentService(studentRepository repository.StudentRepository) StudentService {
	return &studentService{studentRepository}
}

func (s *studentService) FetchAll() ([]model.Student, error) {
	students, err := s.studentRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *studentService) FetchByID(id int) (*model.Student, error) {
	student, err := s.studentRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *studentService) Store(student *model.Student) error {
	err := s.studentRepository.Store(student)
	if err != nil {
		return err
	}

	return nil
}

func (s *studentService) Update(id int, student *model.Student) error {
	err := s.studentRepository.Update(id, student)
	if err != nil {
		return err
	}

	return nil
}

func (s *studentService) Delete(id int) error {
	err := s.studentRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
