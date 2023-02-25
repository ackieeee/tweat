package mysql

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sugartr3e/tweat/domain/entity"
	"gorm.io/gorm"
)

type TweatMysql struct {
	db *sqlx.Tx
}

type TweatGormMysql struct {
	db *gorm.DB
}
type ExecQueryer interface {
	sqlx.Queryer
	sqlx.Execer
}

func NewMysql(db *sqlx.Tx) *TweatMysql {
	return &TweatMysql{db}
}

func NewTweatGormMysql(db *gorm.DB) *TweatGormMysql {
	return &TweatGormMysql{db}
}

func (m *TweatMysql) GetAll(userID string) (entity.TweatLikesList, error) {
	query := strings.Join([]string{
		"SELECT tweats.id, tweats.text, tweats.user_id, count(likes.id) likes_count, users.name user_name",
		"FROM tweats",
		"LEFT JOIN likes ON `tweats`.id=`likes`.tweat_id",
		"LEFT JOIN users ON `tweats`.user_id=`users`.id",
		"WHERE tweats.user_id IN",
		"(SELECT follow_user_id FROM follows WHERE follows.user_id=?)",
		"GROUP BY tweats.id",
	}, " ")
	likesRows, err := m.db.Queryx(query, userID)
	if err != nil {
		return nil, err
	}

	tweats := entity.TweatLikesList{}
	for likesRows.Next() {
		var tweat entity.TweatLikes
		if err := likesRows.StructScan(&tweat); err != nil {
			return nil, err
		}
		tweats = append(tweats, tweat)
	}
	return tweats, nil
}

func (m *TweatGormMysql) GetAll(userID string) (entity.Tweats, error) {
	var tweats []entity.Tweat
	subQuery := m.db.Select("follow_user_id").Where("user_id = ?", userID).Table("follows")
	err := m.db.Debug().
		Preload("Likes").
		Preload("User").
		Where("user_id IN (?) AND parent_id IS NULL", subQuery).
		Find(&tweats).
		Error
	return tweats, err
}

func (m *TweatGormMysql) ExistsLike(tweatID int, userID int) bool {
	like := entity.Like{}
	var exits bool
	err := m.db.Model(&like).
		Select("count(*) > 0").
		Where("tweat_id = ? AND user_id = ?", tweatID, userID).
		Find(&exits).
		Error
	if err != nil {
		return false
	}

	if !exits {
		return false
	}
	return true
}

func (m *TweatGormMysql) AddLike(tweatID int, userID int) error {
	// if exists := m.ExistsLike(tweatID, userID); exists {
	// 	return nil
	// }
	like := entity.Like{
		TweatID: uint(tweatID),
		UserID:  uint(userID),
	}

	return m.db.Create(&like).Error
}

func (m *TweatGormMysql) DeleteLike(tweatID int, userID int) error {
	var like entity.Like
	return m.db.Where("tweat_id = ? AND user_id = ?", tweatID, userID).Delete(&like).Error
}
