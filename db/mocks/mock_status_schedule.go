package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type StatusSchedulesPayload struct {
	StatusSchedules []db.StatusSchedule
	Error           error
}

func (r *MockRepositoryService) ListStatusSchedules(ctx context.Context, evseID int64) ([]db.StatusSchedule, error) {
	if len(r.listStatusSchedulesPayload) == 0 {
		return []db.StatusSchedule{}, nil
	}

	response := r.listStatusSchedulesPayload[0]
	r.listStatusSchedulesPayload = r.listStatusSchedulesPayload[1:]
	return response.StatusSchedules, response.Error
}

func (r *MockRepositoryService) SetListStatusSchedulesPayload(response StatusSchedulesPayload) {
	r.listStatusSchedulesPayload = append(r.listStatusSchedulesPayload, response)
}
