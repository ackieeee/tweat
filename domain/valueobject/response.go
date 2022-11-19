package valueobject

type TweatResponse struct {
	ID       uint   `json:"id"`
	Text     string `json:"text"`
	Likes    []Like `json:"likes"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
}

type Like struct {
	ID      uint `json:"id"`
	TweatID uint `json:"tweat_id"`
	UserID  uint `json:"user_id"`
}
