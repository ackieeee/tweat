package usecase

import (
	"context"
	"github.com/sugartr3e/tweat/repository"

	"github.com/sugartr3e/tweat/domain/entity"
)

type tweatUsecase struct {
	tr repository.TweatRepository
}

type TweatUsecase interface {
	GetAll(ctx context.Context, userID string) (entity.Tweats, error)
	AddLike(ctx context.Context, tweatID int, userID int) error
	DeleteLike(ctx context.Context, tweatID int, userID int) error
	ToggleLike(ctx context.Context, tweatID int, userID int) (bool, error)
}

func NewTweatUsecase(tr repository.TweatRepository) TweatUsecase {
	return &tweatUsecase{tr}
}

func (tu *tweatUsecase) GetAll(ctx context.Context, userID string) (entity.Tweats, error) {
	tweats, err := tu.tr.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tweats, nil
}

func (tu *tweatUsecase) AddLike(ctx context.Context, tweatID int, userID int) error {
	return tu.tr.AddLike(ctx, tweatID, userID)
}

func (tu *tweatUsecase) DeleteLike(ctx context.Context, tweatID int, userID int) error {
	return tu.tr.DeleteLike(ctx, tweatID, userID)
}

func (tu *tweatUsecase) ToggleLike(ctx context.Context, tweatID int, userID int) (bool, error) {
	return tu.tr.ToggleLike(ctx, tweatID, userID)
}
