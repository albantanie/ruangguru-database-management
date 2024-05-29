package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
}

type studentRepoImpl struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	rows, err := s.db.Query("SELECT id, name, address, class FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	row := s.db.QueryRow("SELECT id, name, address, class FROM students WHERE id = $1", id)

	var student model.Student
	err := row.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	query := "INSERT INTO students (name, address, class) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(query, student.Name, student.Address, student.Class)
	return err
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	query := "UPDATE students SET name = $1, address = $2, class = $3 WHERE id = $4"
	_, err := s.db.Exec(query, student.Name, student.Address, student.Class, id)
	return err
}

func (s *studentRepoImpl) Delete(id int) error {
	query := "DELETE FROM students WHERE id = $1"
	_, err := s.db.Exec(query, id)
	return err
}
