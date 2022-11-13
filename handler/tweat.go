package handler

import (
	"errors"
	"net/http"

	"github.com/gba-3/tweat/usecase"
)

type tweatHandler struct {
	tu usecase.TweatUsecase
}

type TweatHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
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

	return http.StatusOK, tweats, nil
}
