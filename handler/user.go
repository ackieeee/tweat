package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gba-3/tweat/domain/entity"
	"github.com/gba-3/tweat/usecase"
)

type userHandler struct {
	uu usecase.UserUsecase
}

type UserHandler interface {
	FindByEmail(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{uu}
}

func (uh *userHandler) FindByEmail(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error(), err
	}
	defer r.Body.Close()

	userBody := entity.User{}
	if err := json.Unmarshal(body, &userBody); err != nil {
		return http.StatusInternalServerError, err.Error(), err
	}
	if userBody.Email == "" {
		return http.StatusBadRequest, err.Error(), err
	}
	if userBody.Password == "" {
		return http.StatusBadRequest, err.Error(), err
	}
	user, err := uh.uu.FindByEmail(userBody.Email, userBody.Password)
	if err != nil {
		return http.StatusBadRequest, err.Error(), err
	}
	return http.StatusOK, user, nil
}
