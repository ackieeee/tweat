package container

import "github.com/sugartr3e/tweat/domain/repository"

func (c Container) GetTweatRepository() repository.TweatRepository {
	return repository.NewTweatRepository()
}

func (c Container) GetUserRepository() repository.UserRepository {
	return repository.NewUserRepository()
}
