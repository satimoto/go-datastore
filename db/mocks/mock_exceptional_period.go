package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ExceptionalPeriodsMockData struct {
	ExceptionalPeriods []db.ExceptionalPeriod
	Error              error
}

func (r *MockRepositoryService) ListExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error) {
	if len(r.listExceptionalOpeningPeriodsMockData) == 0 {
		return []db.ExceptionalPeriod{}, nil
	}

	response := r.listExceptionalOpeningPeriodsMockData[0]
	r.listExceptionalOpeningPeriodsMockData = r.listExceptionalOpeningPeriodsMockData[1:]
	return response.ExceptionalPeriods, response.Error
}

func (r *MockRepositoryService) ListExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error) {
	if len(r.listExceptionalClosingPeriodsMockData) == 0 {
		return []db.ExceptionalPeriod{}, nil
	}

	response := r.listExceptionalClosingPeriodsMockData[0]
	r.listExceptionalClosingPeriodsMockData = r.listExceptionalClosingPeriodsMockData[1:]
	return response.ExceptionalPeriods, response.Error
}

func (r *MockRepositoryService) SetListExceptionalOpeningPeriodsMockData(response ExceptionalPeriodsMockData) {
	r.listExceptionalOpeningPeriodsMockData = append(r.listExceptionalOpeningPeriodsMockData, response)
}

func (r *MockRepositoryService) SetListExceptionalClosingPeriodsMockData(response ExceptionalPeriodsMockData) {
	r.listExceptionalClosingPeriodsMockData = append(r.listExceptionalClosingPeriodsMockData, response)
}
