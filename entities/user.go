package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string
	FirstName string
	Lastname  string
	Role      string
}
