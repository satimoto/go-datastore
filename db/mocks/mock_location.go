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
	if len(r.listLocationsResponse) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsResponse[0]
	r.listLocationsResponse = r.listLocationsResponse[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) SetListLocationsResponse(response LocationsResponse) {
	r.listLocationsResponse = append(r.listLocationsResponse, response)
}
