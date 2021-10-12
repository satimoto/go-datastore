package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type OpeningTimePayload struct {
	OpeningTime db.OpeningTime
	Error       error
}

func (r *MockRepositoryService) CreateOpeningTime(ctx context.Context, twentyfourseven bool) (db.OpeningTime, error) {
	if len(r.createOpeningTimePayload) == 0 {
		return db.OpeningTime{}, ErrorNotFound()
	}

	response := r.createOpeningTimePayload[0]
	r.createOpeningTimePayload = r.createOpeningTimePayload[1:]
	return response.OpeningTime, response.Error
}

func (r *MockRepositoryService) GetOpeningTime(ctx context.Context, id int64) (db.OpeningTime, error) {
	if len(r.getOpeningTimePayload) == 0 {
		return db.OpeningTime{}, ErrorNotFound()
	}

	response := r.createOpeningTimePayload[0]
	r.getOpeningTimePayload = r.getOpeningTimePayload[1:]
	return response.OpeningTime, response.Error
}

func (r *MockRepositoryService) SetCreateOpeningTimePayload(response OpeningTimePayload) {
	r.createOpeningTimePayload = append(r.createOpeningTimePayload, response)
}

func (r *MockRepositoryService) SetGetOpeningTimePayload(response OpeningTimePayload) {
	r.getOpeningTimePayload = append(r.getOpeningTimePayload, response)
}
