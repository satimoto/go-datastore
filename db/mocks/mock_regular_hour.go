package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type RegularHoursPayload struct {
	RegularHours []db.RegularHour
	Error        error
}

func (r *MockRepositoryService) ListRegularHours(ctx context.Context, openingTimeID int64) ([]db.RegularHour, error) {
	if len(r.listRegularHoursPayload) == 0 {
		return []db.RegularHour{}, nil
	}

	response := r.listRegularHoursPayload[0]
	r.listRegularHoursPayload = r.listRegularHoursPayload[1:]
	return response.RegularHours, response.Error
}

func (r *MockRepositoryService) SetListRegularHoursPayload(response RegularHoursPayload) {
	r.listRegularHoursPayload = append(r.listRegularHoursPayload, response)
}
