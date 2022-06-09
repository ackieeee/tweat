package handler

import (
	"fmt"
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
	fmt.Println(ctx.Value("token"))
	tweats, err := th.tu.GetAll(ctx)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, tweats, nil
}
