package repository

import (
	"context"
	"github.com/sugartr3e/tweat/infrastructure/database"

	"github.com/sugartr3e/tweat/domain/entity"
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
	db := database.GetTweatGorm()
	//tweats, err := mysql.NewTweatGormMysql(db).GetAll(userID)
	var tweats []entity.Tweat
	subQuery := db.Select("follow_user_id").Where("user_id = ?", userID).Table("follows")
	err := db.Debug().
		Preload("Likes").
		Preload("User").
		Where("user_id IN (?) AND parent_id IS NULL", subQuery).
		Find(&tweats).
		Error
	if err != nil {
		return nil, err
	}
	return tweats, nil
}

func (tr *tweatRepository) AddLike(ctx context.Context, tweatID int, userID int) error {
	db := database.GetTweatGorm()
	like := entity.Like{
		TweatID: uint(tweatID),
		UserID:  uint(userID),
	}

	tx := db.Begin()
	err := tx.Create(&like).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (tr *tweatRepository) DeleteLike(ctx context.Context, tweatID int, userID int) error {
	db := database.GetTweatGorm()
	var like entity.Like

	tx := db.Begin()
	err := tx.Where("tweat_id = ? AND user_id = ?", tweatID, userID).Delete(&like).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (tr *tweatRepository) ToggleLike(ctx context.Context, tweatID int, userID int) (bool, error) {
	db := database.GetTweatGorm()

	like := entity.Like{}
	var exits bool
	err := db.Model(&like).
		Select("count(*) > 0").
		Where("tweat_id = ? AND user_id = ?", tweatID, userID).
		Find(&exits).
		Error
	if err != nil {
		return false, err
	}
	if exits {
		if err = tr.DeleteLike(ctx, tweatID, userID); err != nil {
			return false, err
		}
		return false, nil
	}
	if err := tr.AddLike(ctx, tweatID, userID); err != nil {
		return false, err
	}
	//if exists := m.ExistsLike(tweatID, userID); exists {
	//	if err := m.DeleteLike(tweatID, userID); err != nil {
	//		return false, err
	//	}
	//	return false, nil
	//}
	//if err := m.AddLike(tweatID, userID); err != nil {
	//	return false, err
	//}
	return true, nil
}
