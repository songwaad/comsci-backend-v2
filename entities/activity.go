package entities

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Name        string
	Description string
}
