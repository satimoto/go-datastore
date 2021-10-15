package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) SetEvseDirection(ctx context.Context, arg db.SetEvseDirectionParams) error {
	r.setEvseDirectionMockData = append(r.setEvseDirectionMockData, arg)
	return nil
}

func (r *MockRepositoryService) ListEvseDirections(ctx context.Context, evseID int64) ([]db.DisplayText, error) {
	if len(r.listEvseDirectionsMockData) == 0 {
		return []db.DisplayText{}, nil
	}

	response := r.listEvseDirectionsMockData[0]
	r.listEvseDirectionsMockData = r.listEvseDirectionsMockData[1:]
	return response.DisplayTexts, response.Error
}

func (r *MockRepositoryService) GetSetEvseDirectionMockData() (db.SetEvseDirectionParams, error) {
	if len(r.setEvseDirectionMockData) == 0 {
		return db.SetEvseDirectionParams{}, ErrorNotFound()
	}

	response := r.setEvseDirectionMockData[0]
	r.setEvseDirectionMockData = r.setEvseDirectionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListEvseDirectionsMockData(response DisplayTextsMockData) {
	r.listEvseDirectionsMockData = append(r.listEvseDirectionsMockData, response)
}
