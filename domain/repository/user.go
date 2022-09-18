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
	CreateUser(name string, email string, password string) error
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
	um := mysql.NewUserMysql(db)
	user, err := um.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	follows, err := um.GetFollowUsers(user.ID)
	if err != nil {
		return nil, err
	}
	user.Follows = follows
	return user, nil
}

func (ur *userRepository) CreateUser(name string, email string, password string) error {
	if name == "" {
		return errors.New("arguments error: name is emtpy.")
	}
	if email == "" {
		return errors.New("arguments error: email is emtpy.")
	}
	if password == "" {
		return errors.New("arguments error: password is emtpy.")
	}
	db := adapter.Tweat()
	return mysql.NewUserMysql(db).CreateUser(name, email, password)
}
