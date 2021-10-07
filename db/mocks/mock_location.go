package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type LocationsResponse struct {
	Locations []db.Location
	Error     error
}

func (r *MockRepositoryService) ListLocations(ctx context.Context) ([]db.Location, error) {
	return r.listLocationsResponse.Locations, r.listLocationsResponse.Error
}

func (r *MockRepositoryService) SetListLocationsResponse(response LocationsResponse) {
	r.listLocationsResponse = response
}
