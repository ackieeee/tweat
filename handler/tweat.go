package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gba-3/tweat/domain/valueobject"
	"github.com/gba-3/tweat/usecase"
)

type tweatHandler struct {
	tu usecase.TweatUsecase
}

type TweatHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
	AddLike(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

func NewTweatHandler(tu usecase.TweatUsecase) TweatHandler {
	return &tweatHandler{tu}
}

func (th *tweatHandler) GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ctx := r.Context()
	uv := ctx.Value("userID")
	userID, ok := uv.(string)
	if !ok {
		return http.StatusInternalServerError, nil, errors.New("Can not get user_id.")
	}
	tweats, err := th.tu.GetAll(ctx, userID)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	resp := []valueobject.TweatResponse{}
	for _, tweat := range tweats {
		t := valueobject.TweatResponse{
			ID:     tweat.ID,
			Text:   tweat.Text,
			Likes:  uint(len(tweat.Likes)),
			UserID: tweat.UserID,
		}
		resp = append(resp, t)
	}

	return http.StatusOK, resp, nil
}

func (th *tweatHandler) AddLike(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ctx := r.Context()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	body := valueobject.AddLikeRequest{}

	if err := json.Unmarshal(buf, &body); err != nil {
		return http.StatusBadRequest, nil, err
	}

	resp := map[string]string{}
	if err := th.tu.AddLike(ctx, body.TweatID, body.UserID); err != nil {
		resp["msg"] = fmt.Sprintf("failed add tweat like. tweat_id:%d, user_id:%d\n", body.TweatID, body.UserID)
		return http.StatusBadRequest, resp, err
	}
	resp["msg"] = "successed."
	return http.StatusOK, resp, nil
}
