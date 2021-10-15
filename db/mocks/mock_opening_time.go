package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type OpeningTimeMockData struct {
	OpeningTime db.OpeningTime
	Error       error
}

func (r *MockRepositoryService) CreateOpeningTime(ctx context.Context, twentyfourseven bool) (db.OpeningTime, error) {
	r.createOpeningTimeMockData = append(r.createOpeningTimeMockData, twentyfourseven)
	return db.OpeningTime{}, nil
}

func (r *MockRepositoryService) GetOpeningTime(ctx context.Context, id int64) (db.OpeningTime, error) {
	if len(r.getOpeningTimeMockData) == 0 {
		return db.OpeningTime{}, ErrorNotFound()
	}

	response := r.getOpeningTimeMockData[0]
	r.getOpeningTimeMockData = r.getOpeningTimeMockData[1:]
	return response.OpeningTime, response.Error
}

func (r *MockRepositoryService) GetCreateOpeningTimeMockData() (bool, error) {
	if len(r.createOpeningTimeMockData) == 0 {
		return false, ErrorNotFound()
	}

	response := r.createOpeningTimeMockData[0]
	r.createOpeningTimeMockData = r.createOpeningTimeMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetOpeningTimeMockData(response OpeningTimeMockData) {
	r.getOpeningTimeMockData = append(r.getOpeningTimeMockData, response)
}
