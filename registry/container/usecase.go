package container

import (
	"github.com/gba-3/tweat/domain/repository"
	"github.com/gba-3/tweat/usecase"
)

func (c Container) GetTweatUsecase(tr repository.TweatRepository) usecase.TweatUsecase {
	return usecase.NewTweatUsecase(tr)
}

func (c Container) GetUserUsecase(tr repository.UserRepository) usecase.UserUsecase {
	return usecase.NewUserUsecase(tr)
}
