package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseImages(ctx context.Context, evseID int64) ([]db.Image, error) {
	if len(r.listEvseImagesPayload) == 0 {
		return []db.Image{}, nil
	}

	response := r.listEvseImagesPayload[0]
	r.listEvseImagesPayload = r.listEvseImagesPayload[1:]
	return response.Images, response.Error
}

func (r *MockRepositoryService) SetListEvseImagesPayload(response ImagesPayload) {
	r.listEvseImagesPayload = append(r.listEvseImagesPayload, response)
}
