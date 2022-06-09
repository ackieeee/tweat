package usecase

import (
	"errors"

	"github.com/gba-3/tweat/domain/entity"
	"github.com/gba-3/tweat/domain/repository"
)

type userUsecase struct {
	ur repository.UserRepository
}

type UserUsecase interface {
	FindByEmail(email string, password string) (*entity.User, error)
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) FindByEmail(email string, password string) (*entity.User, error) {
	if email == "" {
		return nil, errors.New("arguments error: email is emtpy.")
	}
	if password == "" {
		return nil, errors.New("arguments error: password is emtpy.")
	}
	return uu.ur.FindByEmail(email, password)
}
