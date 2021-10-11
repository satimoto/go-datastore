package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationImages(ctx context.Context, locationID int64) ([]db.Image, error) {
	if len(r.listLocationImagesResponse) == 0 {
		return []db.Image{}, nil
	}

	response := r.listLocationImagesResponse[0]
	r.listLocationImagesResponse = r.listLocationImagesResponse[1:]
	return response.Images, response.Error
}

func (r *MockRepositoryService) SetListLocationImagesResponse(response ImagesResponse) {
	r.listLocationImagesResponse = append(r.listLocationImagesResponse, response)
}
