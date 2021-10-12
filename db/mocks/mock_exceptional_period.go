package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ExceptionalPeriodsPayload struct {
	ExceptionalPeriods []db.ExceptionalPeriod
	Error              error
}

func (r *MockRepositoryService) ListExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error) {
	if len(r.listExceptionalOpeningPeriodsPayload) == 0 {
		return []db.ExceptionalPeriod{}, nil
	}

	response := r.listExceptionalOpeningPeriodsPayload[0]
	r.listExceptionalOpeningPeriodsPayload = r.listExceptionalOpeningPeriodsPayload[1:]
	return response.ExceptionalPeriods, response.Error
}

func (r *MockRepositoryService) ListExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error) {
	if len(r.listExceptionalClosingPeriodsPayload) == 0 {
		return []db.ExceptionalPeriod{}, nil
	}

	response := r.listExceptionalClosingPeriodsPayload[0]
	r.listExceptionalClosingPeriodsPayload = r.listExceptionalClosingPeriodsPayload[1:]
	return response.ExceptionalPeriods, response.Error
}

func (r *MockRepositoryService) SetListExceptionalOpeningPeriodsPayload(response ExceptionalPeriodsPayload) {
	r.listExceptionalOpeningPeriodsPayload = append(r.listExceptionalOpeningPeriodsPayload, response)
}

func (r *MockRepositoryService) SetListExceptionalClosingPeriodsPayload(response ExceptionalPeriodsPayload) {
	r.listExceptionalClosingPeriodsPayload = append(r.listExceptionalClosingPeriodsPayload, response)
}
