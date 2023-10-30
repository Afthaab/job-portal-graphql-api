package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required"`
	HashPassword string `json:"hash_password" validate:"required"`
}

type Company struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Salary   string `json:"salary" validate:"required"`
}
