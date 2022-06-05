package entity

type Tweat struct {
	ID     int    `db:"id"`
	Text   string `db:"text"`
	UserID int    `db:"user_id"`
}

type Tweats []Tweat
