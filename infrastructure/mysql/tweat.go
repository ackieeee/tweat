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

func (m *TweatMysql) GetAll() (entity.TweatAlls, error) {
	likesRows, err := m.db.Queryx("SELECT tweats.id, tweats.text, tweats.user_id, likes.tweat_id as likes_tweat_id FROM tweats INNER JOIN likes ON `tweats`.id=`likes`.tweat_id")
	if err != nil {
		return nil, err
	}

	tweats := entity.Tweats{}
	for likesRows.Next() {
		var tweat entity.Tweat
		if err := likesRows.StructScan(&tweat); err != nil {
			return nil, err
		}
		tweats = append(tweats, tweat)
	}
	return tweats.Convert(), nil
}
