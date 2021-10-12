package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ImagePayload struct {
	Image db.Image
	Error error
}

type ImagesPayload struct {
	Images []db.Image
	Error  error
}

func (r *MockRepositoryService) CreateImage(ctx context.Context, arg db.CreateImageParams) (db.Image, error) {
	if len(r.createImagePayload) == 0 {
		return db.Image{}, ErrorNotFound()
	}

	response := r.createImagePayload[0]
	r.createImagePayload = r.createImagePayload[1:]
	return response.Image, response.Error
}

func (r *MockRepositoryService) GetImage(ctx context.Context, id int64) (db.Image, error) {
	if len(r.getImagePayload) == 0 {
		return db.Image{}, ErrorNotFound()
	}

	response := r.getImagePayload[0]
	r.getImagePayload = r.getImagePayload[1:]
	return response.Image, response.Error
}

func (r *MockRepositoryService) SetCreateImagePayload(response ImagePayload) {
	r.createImagePayload = append(r.createImagePayload, response)
}

func (r *MockRepositoryService) SetGetImagePayload(response ImagePayload) {
	r.getImagePayload = append(r.getImagePayload, response)
}
