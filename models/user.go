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
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
