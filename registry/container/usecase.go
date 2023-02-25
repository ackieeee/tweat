package container

import (
	"github.com/sugartr3e/tweat/domain/repository"
	"github.com/sugartr3e/tweat/usecase"
)

func (c Container) GetTweatUsecase(tr repository.TweatRepository) usecase.TweatUsecase {
	return usecase.NewTweatUsecase(tr)
}

func (c Container) GetUserUsecase(tr repository.UserRepository) usecase.UserUsecase {
	return usecase.NewUserUsecase(tr)
}
