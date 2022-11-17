package valueobject

type AddLikeRequest struct {
	TweatID int `json:"tweat_id"`
	UserID  int `json:"user_id"`
}
