package models

import "gorm.io/gorm"

//User example
type User struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name" example:"john"`
	UserName   string `json:"username" example:"johnPork"`
	Password   string `json:"password"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}
