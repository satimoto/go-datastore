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
	if len(r.createImageResponse) == 0 {
		return db.Image{}, ErrorNotFound()
	}

	response := r.createImageResponse[0]
	r.createImageResponse = r.createImageResponse[1:]
	return response.Image, response.Error
}

func (r *MockRepositoryService) GetImage(ctx context.Context, id int64) (db.Image, error) {
	if len(r.getImageResponse) == 0 {
		return db.Image{}, ErrorNotFound()
	}

	response := r.getImageResponse[0]
	r.getImageResponse = r.getImageResponse[1:]
	return response.Image, response.Error
}

func (r *MockRepositoryService) SetCreateImageResponse(response ImageResponse) {
	r.createImageResponse = append(r.createImageResponse, response)
}

func (r *MockRepositoryService) SetGetImageResponse(response ImageResponse) {
	r.getImageResponse = append(r.getImageResponse, response)
}
