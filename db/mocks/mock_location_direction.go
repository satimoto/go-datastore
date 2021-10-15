package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationDirections(ctx context.Context, locationID int64) ([]db.DisplayText, error) {
	if len(r.listLocationDirectionsMockData) == 0 {
		return []db.DisplayText{}, nil
	}

	response := r.listLocationDirectionsMockData[0]
	r.listLocationDirectionsMockData = r.listLocationDirectionsMockData[1:]
	return response.DisplayTexts, response.Error
}

func (r *MockRepositoryService) SetListLocationDirectionsMockData(response DisplayTextsMockData) {
	r.listLocationDirectionsMockData = append(r.listLocationDirectionsMockData, response)
}
