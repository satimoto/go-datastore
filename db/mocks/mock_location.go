package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type LocationResponse struct {
	Location db.Location
	Error    error
}

type LocationsResponse struct {
	Locations []db.Location
	Error     error
}

func (r *MockRepositoryService) GetLocationByUid(ctx context.Context, uid string) (db.Location, error) {
	if len(r.getLocationByUidResponse) == 0 {
		return db.Location{}, nil
	}

	response := r.getLocationByUidResponse[0]
	r.getLocationByUidResponse = r.getLocationByUidResponse[1:]
	return response.Location, response.Error

}

func (r *MockRepositoryService) ListLocations(ctx context.Context) ([]db.Location, error) {
	if len(r.listLocationsResponse) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsResponse[0]
	r.listLocationsResponse = r.listLocationsResponse[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) SetGetLocationByUidResponse(response LocationResponse) {
	r.getLocationByUidResponse = append(r.getLocationByUidResponse, response)
}

func (r *MockRepositoryService) SetListLocationsResponse(response LocationsResponse) {
	r.listLocationsResponse = append(r.listLocationsResponse, response)
}
