package entity

type User struct {
	ID       int     `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	Email    string  `db:"email" json:"email"`
	Password string  `db:"password" json:"password"`
	Follows  Follows `json:"follows"`
}

type Follow struct {
	ID           int `db:"id"`
	UserID       int `db:"user_id"`
	FollowUserID int `db:"follow_user_id"`
}

type Follows []Follow
