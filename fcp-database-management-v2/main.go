package main

import (
	"a21hc3NpZ25tZW50/api"
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"a21hc3NpZ25tZW50/service"
)

func main() {
	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "kampusmerdeka",
		Port:         5432,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.User{}, &model.Session{}, &model.Student{}, &model.Class{})

	classes := []model.Class{
		{
			Name:       "Mathematics",
			Professor:  "Dr. Smith",
			RoomNumber: 101,
		},
		{
			Name:       "Physics",
			Professor:  "Dr. Johnson",
			RoomNumber: 102,
		},
		{
			Name:       "Chemistry",
			Professor:  "Dr. Lee",
			RoomNumber: 103,
		},
	}

	for _, c := range classes {
		if err := conn.Create(&c).Error; err != nil {
			panic("failed to create default classes")
		}
	}

	userRepo := repo.NewUserRepo(conn)
	sessionRepo := repo.NewSessionRepo(conn)
	studentRepo := repo.NewStudentRepo(conn)
	classRepo := repo.NewClassRepo(conn)

	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	studentService := service.NewStudentService(studentRepo)
	classService := service.NewClassService(classRepo)

	mainAPI := api.NewAPI(userService, sessionService, studentService, classService)
	mainAPI.Start()
}
