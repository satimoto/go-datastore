package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseImages(ctx context.Context, evseID int64) ([]db.Image, error) {
	if len(r.listEvseImagesResponse) == 0 {
		return []db.Image{}, nil
	}

	response := r.listEvseImagesResponse[0]
	r.listEvseImagesResponse = r.listEvseImagesResponse[1:]
	return response.Images, response.Error
}

func (r *MockRepositoryService) SetListEvseImagesResponse(response ImagesResponse) {
	r.listEvseImagesResponse = append(r.listEvseImagesResponse, response)
}
