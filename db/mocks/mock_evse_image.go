package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseImages(ctx context.Context, evseID int64) ([]db.Image, error) {
	return r.listEvseImagesResponse.Images, r.listEvseImagesResponse.Error
}

func (r *MockRepositoryService) SetListEvseImagesResponse(response ImagesResponse) {
	r.listEvseImagesResponse = response
}
