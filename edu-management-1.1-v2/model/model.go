package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type Student struct {
	ID      int
	Name    string
	Address string
	Class   string
}

type Teacher struct {
	ID      int
	Name    string
	Address string
	Subject string
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}
