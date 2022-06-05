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

func (m *TweatMysql) GetAll() (entity.Tweats, error) {
	rows, err := m.db.Queryx("SELECT id, text, user_id FROM tweats")
	if err != nil {
		return nil, err
	}

	tweats := entity.Tweats{}
	for rows.Next() {
		var tweat entity.Tweat
		if err := rows.StructScan(&tweat); err != nil {
			return nil, err
		}
		tweats = append(tweats, tweat)
	}
	return tweats, nil
}
