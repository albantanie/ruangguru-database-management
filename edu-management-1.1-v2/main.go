package main

import (
	"a21hc3NpZ25tZW50/api"
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"database/sql"
	"fmt"
	"log"

	_ "embed"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(creds *model.Credential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func SQLExecute(db *sql.DB) error {
	//create table students
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS students (id SERIAL PRIMARY KEY, name VARCHAR(255), address VARCHAR(255), class VARCHAR(255))")
	if err != nil {
		return err
	}

	//create table teachers
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS teachers (id SERIAL PRIMARY KEY, name VARCHAR(255), address VARCHAR(255), subject VARCHAR(255))")
	if err != nil {
		return err
	}

	return nil
}

func Reset(db *sql.DB, table string) error {
	_, err := db.Exec("TRUNCATE " + table)
	if err != nil {
		return err
	}

	_, err = db.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "kampusmerdeka",
		Port:         5432,
	}
	dbConn, err := Connect(&dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	err = SQLExecute(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	studentRepo := repo.NewStudentRepo(dbConn)
	teacherRepo := repo.NewTeacherRepo(dbConn)

	mainAPI := api.NewAPI(studentRepo, teacherRepo)
	mainAPI.Start()
}
