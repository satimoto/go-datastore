package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationImages(ctx context.Context, locationID int64) ([]db.Image, error) {
	if len(r.listLocationImagesPayload) == 0 {
		return []db.Image{}, nil
	}

	response := r.listLocationImagesPayload[0]
	r.listLocationImagesPayload = r.listLocationImagesPayload[1:]
	return response.Images, response.Error
}

func (r *MockRepositoryService) SetListLocationImagesPayload(response ImagesPayload) {
	r.listLocationImagesPayload = append(r.listLocationImagesPayload, response)
}
