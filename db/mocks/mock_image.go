package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ImageResponse struct {
	Image db.Image
	Error error
}

type ImagesResponse struct {
	Images []db.Image
	Error  error
}

func (r *MockRepositoryService) CreateImage(ctx context.Context, arg db.CreateImageParams) (db.Image, error) {
	return r.createImageResponse.Image, r.createImageResponse.Error
}

func (r *MockRepositoryService) GetImage(ctx context.Context, id int64) (db.Image, error) {
	return r.getImageResponse.Image, r.getImageResponse.Error
}

func (r *MockRepositoryService) SetCreateImageResponse(response ImageResponse) {
	r.createImageResponse = response
}

func (r *MockRepositoryService) SetGetImageResponse(response ImageResponse) {
	r.getImageResponse = response
}
