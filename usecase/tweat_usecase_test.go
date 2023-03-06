package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_repository "github.com/sugartr3e/tweat/mock"
	"testing"

	"github.com/sugartr3e/tweat/domain/entity"
)

type MockTweatRepository struct {
	getAll func(userID string) (entity.TweatLikesList, error)
}

func (tr *MockTweatRepository) GetAll(ctx context.Context, userID string) (entity.TweatLikesList, error) {
	return tr.getAll(userID)
}

func TestGetAll(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tweatRepositoryMock := mock_repository.NewMockTweatRepository(ctrl)

	ctx := context.Background()
	tweats := entity.Tweats{
		{
			ID:     1,
			Text:   "test",
			UserID: 1,
		},
	}
	tweatRepositoryMock.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(tweats, nil)

	testCase := struct {
		expectList  entity.Tweats
		expectError error
	}{
		entity.Tweats{
			{
				ID:     1,
				Text:   "test",
				UserID: 1,
			},
		},
		nil,
	}
	tu := NewTweatUsecase(tweatRepositoryMock)
	list, err := tu.GetAll(ctx, "1")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(t, testCase.expectList, list)
}
