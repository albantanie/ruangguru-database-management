package api

import (
	"a21hc3NpZ25tZW50/service"
	"fmt"
	"net/http"
)

type API struct {
	userService    service.UserService
	sessionService service.SessionService
	studentService service.StudentService
	mux            *http.ServeMux
}

func NewAPI(userService service.UserService, sessionService service.SessionService, studentService service.StudentService) API {
	mux := http.NewServeMux()
	api := API{
		userService,
		sessionService,
		studentService,
		mux,
	}

	mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))
	mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))
	mux.Handle("/user/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))

	mux.Handle("/student/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllStudent))))
	mux.Handle("/student/get", api.Get(api.Auth(http.HandlerFunc(api.FetchStudentByID))))
	mux.Handle("/student/add", api.Post(api.Auth(http.HandlerFunc(api.Storestudent))))
	mux.Handle("/student/update", api.Put(api.Auth(http.HandlerFunc(api.Updatestudent))))
	mux.Handle("/student/delete", api.Delete(http.HandlerFunc(api.Deletestudent)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
