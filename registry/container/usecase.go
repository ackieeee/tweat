package container

import (
	repository2 "github.com/sugartr3e/tweat/repository"
	"github.com/sugartr3e/tweat/usecase"
)

func (c Container) GetTweatUsecase(tr repository2.TweatRepository) usecase.TweatUsecase {
	return usecase.NewTweatUsecase(tr)
}

func (c Container) GetUserUsecase(tr repository2.UserRepository) usecase.UserUsecase {
	return usecase.NewUserUsecase(tr)
}
