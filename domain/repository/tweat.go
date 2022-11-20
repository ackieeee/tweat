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
	AddLike(ctx context.Context, tweatID int, userID int) error
	DeleteLike(ctx context.Context, tweatID int, userID int) error
	ToggleLike(ctx context.Context, tweatID int, userID int) (bool, error)
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

func (tr *tweatRepository) AddLike(ctx context.Context, tweatID int, userID int) error {
	db := adapter.TweatGorm()
	return mysql.NewTweatGormMysql(db).AddLike(tweatID, userID)
}

func (tr *tweatRepository) DeleteLike(ctx context.Context, tweatID int, userID int) error {
	db := adapter.TweatGorm()
	return mysql.NewTweatGormMysql(db).DeleteLike(tweatID, userID)
}

func (tr *tweatRepository) ToggleLike(ctx context.Context, tweatID int, userID int) (bool, error) {
	db := adapter.TweatGorm()
	m := mysql.NewTweatGormMysql(db)
	if exists := m.ExistsLike(tweatID, userID); exists {
		if err := m.DeleteLike(tweatID, userID); err != nil {
			return false, err
		}
		return false, nil
	}
	if err := m.AddLike(tweatID, userID); err != nil {
		return false, err
	}
	return true, nil
}
