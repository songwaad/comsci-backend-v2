package usecases

import (
	"github.com/songwaad/cs-activity-backend/entities"
)

type ActivityRepo interface {
	Create(activity *entities.Activity) error
	FindAll() ([]entities.Activity, error)
	FindById(id uint) (*entities.Activity, error)
	UpdateById(id uint, activity *entities.Activity) error
	DeleteById(id uint) error
}
