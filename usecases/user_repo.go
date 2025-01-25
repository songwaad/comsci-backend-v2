package usecases

import (
	"github.com/songwaad/cs-activity-backend/entities"
)

type UserRepo interface {
	Create(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
}
