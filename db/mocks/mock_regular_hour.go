package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type RegularHoursResponse struct {
	RegularHours []db.RegularHour
	Error        error
}

func (r *MockRepositoryService) ListRegularHours(ctx context.Context, openingTimeID int64) ([]db.RegularHour, error) {
	if len(r.listRegularHoursResponse) == 0 {
		return []db.RegularHour{}, nil
	}

	response := r.listRegularHoursResponse[0]
	r.listRegularHoursResponse = r.listRegularHoursResponse[1:]
	return response.RegularHours, response.Error
}

func (r *MockRepositoryService) SetListRegularHoursResponse(response RegularHoursResponse) {
	r.listRegularHoursResponse = append(r.listRegularHoursResponse, response)
}
