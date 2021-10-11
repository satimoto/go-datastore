package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type OpeningTimeResponse struct {
	OpeningTime db.OpeningTime
	Error       error
}

func (r *MockRepositoryService) CreateOpeningTime(ctx context.Context, twentyfourseven bool) (db.OpeningTime, error) {
	if len(r.createOpeningTimeResponse) == 0 {
		return db.OpeningTime{}, ErrorNotFound()
	}

	response := r.createOpeningTimeResponse[0]
	r.createOpeningTimeResponse = r.createOpeningTimeResponse[1:]
	return response.OpeningTime, response.Error
}

func (r *MockRepositoryService) GetOpeningTime(ctx context.Context, id int64) (db.OpeningTime, error) {
	if len(r.getOpeningTimeResponse) == 0 {
		return db.OpeningTime{}, ErrorNotFound()
	}

	response := r.createOpeningTimeResponse[0]
	r.getOpeningTimeResponse = r.getOpeningTimeResponse[1:]
	return response.OpeningTime, response.Error
}

func (r *MockRepositoryService) SetCreateOpeningTimeResponse(response OpeningTimeResponse) {
	r.createOpeningTimeResponse = append(r.createOpeningTimeResponse, response)
}

func (r *MockRepositoryService) SetGetOpeningTimeResponse(response OpeningTimeResponse) {
	r.getOpeningTimeResponse = append(r.getOpeningTimeResponse, response)
}
