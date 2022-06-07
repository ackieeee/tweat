package entity

type TweatLikes struct {
	ID     int    `db:"id"`
	Text   string `db:"text"`
	UserID int    `db:"user_id"`
	Likes  int    `db:"likes_count"`
}

type TweatLikesList []TweatLikes

type Tweat struct {
	ID           int    `db:"id"`
	Text         string `db:"text"`
	TweatUserID  int    `db:"user_id"`
	LikesTweatID int    `db:"likes_tweat_id"`
}

type Tweats []Tweat

func (ts Tweats) Convert() TweatAlls {
	var tweats TweatAlls
	likes := map[int]int{}
	for _, t := range ts {
		if count, ok := likes[t.LikesTweatID]; ok {
			likes[t.LikesTweatID] = count + 1
		} else {
			likes[t.LikesTweatID] = 1
			tweat := TweatAll{
				ID:     t.ID,
				Text:   t.Text,
				UserID: t.TweatUserID,
				Likes:  1,
			}
			tweats = append(tweats, &tweat)
		}
	}

	for _, tweat := range tweats {
		tweat.Likes = likes[tweat.ID]
	}
	return tweats
}

type TweatAll struct {
	ID     int
	Text   string
	UserID int
	Likes  int
}

type TweatAlls []*TweatAll
