package usecase

import (
	"context"

	"github.com/gba-3/tweat/domain/entity"
	"github.com/gba-3/tweat/domain/repository"
)

type tweatUsecase struct {
	tr repository.TweatRepository
}

type TweatUsecase interface {
	GetAll(ctx context.Context) (entity.TweatAlls, error)
}

func NewTweatUsecase(tr repository.TweatRepository) TweatUsecase {
	return &tweatUsecase{tr}
}

func (tu *tweatUsecase) GetAll(ctx context.Context) (entity.TweatAlls, error) {
	tweats, err := tu.tr.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tweats, nil
}
