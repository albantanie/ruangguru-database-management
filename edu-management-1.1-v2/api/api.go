package api

import (
	repo "a21hc3NpZ25tZW50/repository"
	"fmt"
	"net/http"
)

type API struct {
	studentRepo repo.StudentRepository
	teacherRepo repo.TeacherRepository
	mux         *http.ServeMux
}

func NewAPI(studentRepo repo.StudentRepository, teacherRepo repo.TeacherRepository) API {
	mux := http.NewServeMux()
	api := API{
		studentRepo,
		teacherRepo,
		mux,
	}

	mux.Handle("/student/get-all", api.Get(http.HandlerFunc(api.FetchAllStudent)))
	mux.Handle("/student/get", api.Get(http.HandlerFunc(api.FetchStudentByID)))
	mux.Handle("/student/add", api.Post(http.HandlerFunc(api.Storestudent)))

	mux.Handle("/teacher/get-all", api.Get(http.HandlerFunc(api.FetchAllTeacher)))
	mux.Handle("/teacher/get", api.Get(http.HandlerFunc(api.FetchTeacherByID)))
	mux.Handle("/teacher/add", api.Post(http.HandlerFunc(api.StoreTeacher)))
	mux.Handle("/teacher/update", api.Put(http.HandlerFunc(api.UpdateTeacher)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
