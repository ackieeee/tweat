package repository

import (
	"context"

	"github.com/gba-3/tweat/domain/entity"
	"github.com/gba-3/tweat/infrastructure/mysql"
	"github.com/gba-3/tweat/infrastructure/mysql/adapter"
)

type tweatRepository struct {
}

type TweatRepository interface {
	GetAll(ctx context.Context) (entity.TweatAlls, error)
}

func NewTweatRepository() TweatRepository {
	return &tweatRepository{}
}

func (tr *tweatRepository) GetAll(ctx context.Context) (entity.TweatAlls, error) {
	db := adapter.Tweat()
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	tweats, err := mysql.NewMysql(tx).GetAll()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return tweats, nil
}
