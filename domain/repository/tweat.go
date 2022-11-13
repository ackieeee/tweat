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
	GetAll(ctx context.Context, userID string) (entity.Tweats, error)
}

func NewTweatRepository() TweatRepository {
	return &tweatRepository{}
}

func (tr *tweatRepository) GetAll(ctx context.Context, userID string) (entity.Tweats, error) {
	// db := adapter.Tweat()
	// tx, err := db.Beginx()
	// if err != nil {
	// 	return nil, err
	// }
	// tweats, err := mysql.NewMysql(tx).GetAll(userID)
	// if err != nil {
	// 	err := tx.Rollback()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return nil, err
	// }

	// if err := tx.Commit(); err != nil {
	// 	err := tx.Rollback()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return nil, err
	// }
	db := adapter.TweatGorm()
	tweats, err := mysql.NewTweatGormMysql(db).GetAll(userID)
	if err != nil {
		return nil, err
	}
	return tweats, nil
}
