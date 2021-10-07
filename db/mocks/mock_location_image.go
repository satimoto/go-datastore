package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationImages(ctx context.Context, locationID int64) ([]db.Image, error) {
	return r.listLocationImagesResponse.Images, r.listLocationImagesResponse.Error
}

func (r *MockRepositoryService) SetListLocationImagesResponse(response ImagesResponse) {
	r.listLocationImagesResponse = response
}
