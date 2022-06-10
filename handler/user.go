package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gba-3/tweat/domain/entity"
	"github.com/gba-3/tweat/usecase"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	uu usecase.UserUsecase
}

type UserHandler interface {
	Signup(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
	Login(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{uu}
}

func (uh *userHandler) Signup(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	return http.StatusBadRequest, err.Error(), err
	// }
	defer r.Body.Close()

	userBody := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		return http.StatusInternalServerError, err.Error(), err
	}

	// GenerateFromPasswordでパスワードをハッシュ化
	// 第2引数はコスト
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), 10)
	if err != nil {
		return http.StatusInternalServerError, err.Error(), err
	}
	if err := uh.uu.CreateUser(userBody.Name, userBody.Email, string(hashedPassword)); err != nil {
		return http.StatusBadRequest, err.Error(), err
	}

	res := map[string]string{
		"message": "successed create user.",
	}
	return http.StatusCreated, res, nil
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
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

	// jwt token
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.ID),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return http.StatusInternalServerError, err.Error(), err
	}
	res := map[string]string{
		"token": token,
	}
	return http.StatusOK, res, nil
}
