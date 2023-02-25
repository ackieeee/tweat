package container

import (
	"github.com/sugartr3e/tweat/handler"
	"github.com/sugartr3e/tweat/usecase"
)

func (c Container) GetTweatHandler(tu usecase.TweatUsecase) handler.TweatHandler {
	return handler.NewTweatHandler(tu)
}

func (c Container) GetUserHandler(tu usecase.UserUsecase) handler.UserHandler {
	return handler.NewUserHandler(tu)
}
