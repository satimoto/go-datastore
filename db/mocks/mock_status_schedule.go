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
	if len(r.listStatusSchedulesResponse) == 0 {
		return []db.StatusSchedule{}, nil
	}

	response := r.listStatusSchedulesResponse[0]
	r.listStatusSchedulesResponse = r.listStatusSchedulesResponse[1:]
	return response.StatusSchedules, response.Error
}

func (r *MockRepositoryService) SetListStatusSchedulesResponse(response StatusSchedulesResponse) {
	r.listStatusSchedulesResponse = append(r.listStatusSchedulesResponse, response)
}
