package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationFacilities(ctx context.Context, locationID int64) ([]db.Facility, error) {
	if len(r.listLocationFacilitiesResponse) == 0 {
		return []db.Facility{}, nil
	}

	response := r.listLocationFacilitiesResponse[0]
	r.listLocationFacilitiesResponse = r.listLocationFacilitiesResponse[1:]
	return response.Facilities, response.Error
}

func (r *MockRepositoryService) SetListLocationFacilitiesResponse(response FacilitiesResponse) {
	r.listLocationFacilitiesResponse = append(r.listLocationFacilitiesResponse, response)
}
