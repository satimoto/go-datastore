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
	return r.listRegularHoursResponse.RegularHours, r.listRegularHoursResponse.Error
}

func (r *MockRepositoryService) SetListRegularHoursResponse(response RegularHoursResponse) {
	r.listRegularHoursResponse = response
}
