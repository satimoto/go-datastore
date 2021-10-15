package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationImages(ctx context.Context, locationID int64) ([]db.Image, error) {
	if len(r.listLocationImagesMockData) == 0 {
		return []db.Image{}, nil
	}

	response := r.listLocationImagesMockData[0]
	r.listLocationImagesMockData = r.listLocationImagesMockData[1:]
	return response.Images, response.Error
}

func (r *MockRepositoryService) SetListLocationImagesMockData(response ImagesMockData) {
	r.listLocationImagesMockData = append(r.listLocationImagesMockData, response)
}
