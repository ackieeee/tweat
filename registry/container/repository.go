package container

import "github.com/gba-3/tweat/domain/repository"

func (c Container) GetTweatRepository() repository.TweatRepository {
	return repository.NewTweatRepository()
}
