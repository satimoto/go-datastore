package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationFacilities(ctx context.Context, locationID int64) ([]db.Facility, error) {
	return r.listLocationFacilitiesResponse.Facilities, r.listLocationFacilitiesResponse.Error
}

func (r *MockRepositoryService) SetListLocationFacilitiesResponse(response FacilitiesResponse) {
	r.listLocationFacilitiesResponse = response
}
