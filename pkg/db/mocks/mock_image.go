package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type ImageMockData struct {
	Image db.Image
	Error error
}

type ImagesMockData struct {
	Images []db.Image
	Error  error
}

func (r *MockRepositoryService) CreateImage(ctx context.Context, arg db.CreateImageParams) (db.Image, error) {
	r.createImageMockData = append(r.createImageMockData, arg)
	return db.Image{}, nil
}

func (r *MockRepositoryService) GetImage(ctx context.Context, id int64) (db.Image, error) {
	if len(r.getImageMockData) == 0 {
		return db.Image{}, ErrorNotFound()
	}

	response := r.getImageMockData[0]
	r.getImageMockData = r.getImageMockData[1:]
	return response.Image, response.Error
}

func (r *MockRepositoryService) GetCreateImageMockData() (db.CreateImageParams, error) {
	if len(r.createImageMockData) == 0 {
		return db.CreateImageParams{}, ErrorNotFound()
	}

	response := r.createImageMockData[0]
	r.createImageMockData = r.createImageMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetImageMockData(response ImageMockData) {
	r.getImageMockData = append(r.getImageMockData, response)
}
