package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationFacilities(ctx context.Context, locationID int64) ([]db.Facility, error) {
	if len(r.listLocationFacilitiesMockData) == 0 {
		return []db.Facility{}, nil
	}

	response := r.listLocationFacilitiesMockData[0]
	r.listLocationFacilitiesMockData = r.listLocationFacilitiesMockData[1:]
	return response.Facilities, response.Error
}

func (r *MockRepositoryService) SetListLocationFacilitiesMockData(response FacilitiesMockData) {
	r.listLocationFacilitiesMockData = append(r.listLocationFacilitiesMockData, response)
}
