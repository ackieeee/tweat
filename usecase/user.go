package usecase

import (
	"errors"

	"github.com/sugartr3e/tweat/domain/entity"
	"github.com/sugartr3e/tweat/domain/repository"
)

type userUsecase struct {
	ur repository.UserRepository
}

type UserUsecase interface {
	FindByEmail(email string, password string) (*entity.User, error)
	CreateUser(name string, email string, password string) error
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

func (uu *userUsecase) CreateUser(name string, email string, password string) error {
	if name == "" {
		return errors.New("arguments error: name is emtpy.")
	}
	if email == "" {
		return errors.New("arguments error: email is emtpy.")
	}
	if password == "" {
		return errors.New("arguments error: password is emtpy.")
	}
	return uu.ur.CreateUser(name, email, password)
}
