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
	return r.listExceptionalOpeningPeriodsResponse.ExceptionalPeriods, r.listExceptionalOpeningPeriodsResponse.Error
}

func (r *MockRepositoryService) ListExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error) {
	return r.listExceptionalClosingPeriodsResponse.ExceptionalPeriods, r.listExceptionalClosingPeriodsResponse.Error
}

func (r *MockRepositoryService) SetListExceptionalOpeningPeriodsResponse(response ExceptionalPeriodsResponse) {
	r.listExceptionalOpeningPeriodsResponse = response
}

func (r *MockRepositoryService) SetListExceptionalClosingPeriodsResponse(response ExceptionalPeriodsResponse) {
	r.listExceptionalClosingPeriodsResponse = response
}
