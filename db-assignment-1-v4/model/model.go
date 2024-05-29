package model

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `gorm:"type:varchar(100);unique"`
	Password string `json:"password"`
}

type Session struct {
	ID       int       `json:"id"`
	Token    string    `json:"token"`
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Class   string `json:"class"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}
