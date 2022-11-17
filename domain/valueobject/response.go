package valueobject

type TweatResponse struct {
	ID     uint   `json:"id"`
	Text   string `json:"text"`
	Likes  uint   `json:"likes"`
	UserID uint   `json:"user_id"`
}
