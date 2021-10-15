package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) SetEvseImage(ctx context.Context, arg db.SetEvseImageParams) error {
	r.setEvseImageMockData = append(r.setEvseImageMockData, arg)
	return nil
}

func (r *MockRepositoryService) ListEvseImages(ctx context.Context, evseID int64) ([]db.Image, error) {
	if len(r.listEvseImagesMockData) == 0 {
		return []db.Image{}, nil
	}

	response := r.listEvseImagesMockData[0]
	r.listEvseImagesMockData = r.listEvseImagesMockData[1:]
	return response.Images, response.Error
}

func (r *MockRepositoryService) GetSetEvseImageMockData() (db.SetEvseImageParams, error) {
	if len(r.setEvseImageMockData) == 0 {
		return db.SetEvseImageParams{}, ErrorNotFound()
	}

	response := r.setEvseImageMockData[0]
	r.setEvseImageMockData = r.setEvseImageMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListEvseImagesMockData(response ImagesMockData) {
	r.listEvseImagesMockData = append(r.listEvseImagesMockData, response)
}
