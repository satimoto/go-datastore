package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ExceptionalPeriodsResponse struct {
	ExceptionalPeriods []db.ExceptionalPeriod
	Error              error
}

func (r *MockRepositoryService) ListExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error) {
	if len(r.listExceptionalOpeningPeriodsResponse) == 0 {
		return []db.ExceptionalPeriod{}, nil
	}

	response := r.listExceptionalOpeningPeriodsResponse[0]
	r.listExceptionalOpeningPeriodsResponse = r.listExceptionalOpeningPeriodsResponse[1:]
	return response.ExceptionalPeriods, response.Error
}

func (r *MockRepositoryService) ListExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error) {
	if len(r.listExceptionalClosingPeriodsResponse) == 0 {
		return []db.ExceptionalPeriod{}, nil
	}

	response := r.listExceptionalClosingPeriodsResponse[0]
	r.listExceptionalClosingPeriodsResponse = r.listExceptionalClosingPeriodsResponse[1:]
	return response.ExceptionalPeriods, response.Error
}

func (r *MockRepositoryService) SetListExceptionalOpeningPeriodsResponse(response ExceptionalPeriodsResponse) {
	r.listExceptionalOpeningPeriodsResponse = append(r.listExceptionalOpeningPeriodsResponse, response)
}

func (r *MockRepositoryService) SetListExceptionalClosingPeriodsResponse(response ExceptionalPeriodsResponse) {
	r.listExceptionalClosingPeriodsResponse = append(r.listExceptionalClosingPeriodsResponse, response)
}
