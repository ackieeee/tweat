package usecase

import (
	"context"
	"reflect"
	"strconv"
	"testing"

	"github.com/gba-3/tweat/domain/entity"
)

type MockTweatRepository struct {
	getAll func(userID string) (entity.TweatLikesList, error)
}

func (tr *MockTweatRepository) GetAll(ctx context.Context, userID string) (entity.TweatLikesList, error) {
	return tr.getAll(userID)
}

func TestGetAll(t *testing.T) {
	testCase := struct {
		expectList  entity.TweatLikesList
		expectError error
	}{
		entity.TweatLikesList{
			{
				ID:     1,
				Text:   "test",
				UserID: 1,
				Likes:  3,
			},
		},
		nil,
	}
	tr := &MockTweatRepository{
		getAll: func(userID string) (entity.TweatLikesList, error) {
			i, err := strconv.Atoi(userID)
			if err != nil {
				return nil, err
			}
			return entity.TweatLikesList{
				{
					ID:     1,
					Text:   "test",
					UserID: i,
					Likes:  3,
				},
			}, nil
		},
	}

	tu := NewTweatUsecase(tr)
	ctx := context.Background()
	list, err := tu.GetAll(ctx, "1")
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(testCase.expectList, list) {
		t.Fatalf("unexpected result. expected=%v actual=%v\n", testCase.expectList, list)
	}
}
