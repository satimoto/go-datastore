package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type StatusSchedulesResponse struct {
	StatusSchedules []db.StatusSchedule
	Error           error
}

func (r *MockRepositoryService) ListStatusSchedules(ctx context.Context, evseID int64) ([]db.StatusSchedule, error) {
	return r.listStatusSchedulesResponse.StatusSchedules, r.listStatusSchedulesResponse.Error
}

func (r *MockRepositoryService) SetListStatusSchedulesResponse(response StatusSchedulesResponse) {
	r.listStatusSchedulesResponse = response
}
