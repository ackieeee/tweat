package repository

import (
	"errors"

	"github.com/gba-3/tweat/domain/entity"
	"github.com/gba-3/tweat/infrastructure/mysql"
	"github.com/gba-3/tweat/infrastructure/mysql/adapter"
)

type userRepository struct {
}

type UserRepository interface {
	FindByEmail(email string, password string) (*entity.User, error)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) FindByEmail(email string, password string) (*entity.User, error) {
	if email == "" {
		return nil, errors.New("arguments error: email is emtpy.")
	}
	if password == "" {
		return nil, errors.New("arguments error: password is emtpy.")
	}

	db := adapter.Tweat()
	return mysql.NewUserMysql(db).FindByEmail(email)
}
