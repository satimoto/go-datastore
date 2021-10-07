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
	return r.createOpeningTimeResponse.OpeningTime, r.createOpeningTimeResponse.Error
}

func (r *MockRepositoryService) GetOpeningTime(ctx context.Context, id int64) (db.OpeningTime, error) {
	return r.getOpeningTimeResponse.OpeningTime, r.getOpeningTimeResponse.Error
}

func (r *MockRepositoryService) SetCreateOpeningTimeResponse(response OpeningTimeResponse) {
	r.createOpeningTimeResponse = response
}

func (r *MockRepositoryService) SetGetOpeningTimeResponse(response OpeningTimeResponse) {
	r.getOpeningTimeResponse = response
}
