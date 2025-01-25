package usecases

import (
	"github.com/songwaad/cs-activity-backend/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepo UserRepo
}

func (u *UserUseCase) Register(user entities.User) error {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return u.UserRepo.Create(&user)
}

func (u *UserUseCase) Login(email, password string) (*entities.User, error) {
	user, err := u.UserRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
