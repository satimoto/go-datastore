package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type RegularHoursMockData struct {
	RegularHours []db.RegularHour
	Error        error
}

func (r *MockRepositoryService) ListRegularHours(ctx context.Context, openingTimeID int64) ([]db.RegularHour, error) {
	if len(r.listRegularHoursMockData) == 0 {
		return []db.RegularHour{}, nil
	}

	response := r.listRegularHoursMockData[0]
	r.listRegularHoursMockData = r.listRegularHoursMockData[1:]
	return response.RegularHours, response.Error
}

func (r *MockRepositoryService) SetListRegularHoursMockData(response RegularHoursMockData) {
	r.listRegularHoursMockData = append(r.listRegularHoursMockData, response)
}
