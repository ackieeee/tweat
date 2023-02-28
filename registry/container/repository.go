package container

import (
	repository2 "github.com/sugartr3e/tweat/repository"
)

func (c Container) GetTweatRepository() repository2.TweatRepository {
	return repository2.NewTweatRepository()
}

func (c Container) GetUserRepository() repository2.UserRepository {
	return repository2.NewUserRepository()
}
