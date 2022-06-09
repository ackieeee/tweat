package mysql

import (
	"github.com/gba-3/tweat/domain/entity"
	"github.com/jmoiron/sqlx"
)

type UserMysql struct {
	db *sqlx.DB
}

func NewUserMysql(db *sqlx.DB) *UserMysql {
	return &UserMysql{db}
}

func (um *UserMysql) FindByEmail(email string) (*entity.User, error) {
	query := "SELECT id, email, password FROM users WHERE email=?"

	user := entity.User{}
	if err := um.db.Get(&user, query, email); err != nil {
		return nil, err
	}
	return &user, nil
}
