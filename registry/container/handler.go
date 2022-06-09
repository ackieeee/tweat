package container

import (
	"github.com/gba-3/tweat/handler"
	"github.com/gba-3/tweat/usecase"
)

func (c Container) GetTweatHandler(tu usecase.TweatUsecase) handler.TweatHandler {
	return handler.NewTweatHandler(tu)
}

func (c Container) GetUserHandler(tu usecase.UserUsecase) handler.UserHandler {
	return handler.NewUserHandler(tu)
}
