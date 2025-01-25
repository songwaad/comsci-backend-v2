package adapters

import (
	"github.com/songwaad/cs-activity-backend/entities"
	"gorm.io/gorm"
)

type GormUserRepo struct {
	DB *gorm.DB
}

func (r *GormUserRepo) Create(user *entities.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepo) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
