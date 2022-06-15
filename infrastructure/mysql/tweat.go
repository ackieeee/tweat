package mysql

import (
	"github.com/gba-3/tweat/domain/entity"
	"github.com/jmoiron/sqlx"
)

type TweatMysql struct {
	db *sqlx.Tx
}
type ExecQueryer interface {
	sqlx.Queryer
	sqlx.Execer
}

func NewMysql(db *sqlx.Tx) *TweatMysql {
	return &TweatMysql{db}
}

func (m *TweatMysql) GetAll(userID string) (entity.TweatLikesList, error) {
	likesRows, err := m.db.Queryx("SELECT tweats.id, tweats.text, tweats.user_id, count(likes.id) likes_count FROM tweats INNER JOIN likes ON `tweats`.id=`likes`.tweat_id WHERE `tweats`.user_id=? GROUP BY tweats.id", userID)
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
