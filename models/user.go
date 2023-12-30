package models

import "gorm.io/gorm"

type Model struct {
	gorm.Model
}

type User struct {
	Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
