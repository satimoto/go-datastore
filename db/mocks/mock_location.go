package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type LocationPayload struct {
	Location db.Location
	Error    error
}

type LocationsPayload struct {
	Locations []db.Location
	Error     error
}

func (r *MockRepositoryService) GetLocation(ctx context.Context, id int64) (db.Location, error) {
	if len(r.getLocationPayload) == 0 {
		return db.Location{}, nil
	}

	response := r.getLocationPayload[0]
	r.getLocationPayload = r.getLocationPayload[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) GetLocationByUid(ctx context.Context, uid string) (db.Location, error) {
	if len(r.getLocationByUidPayload) == 0 {
		return db.Location{}, nil
	}

	response := r.getLocationByUidPayload[0]
	r.getLocationByUidPayload = r.getLocationByUidPayload[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) ListLocations(ctx context.Context) ([]db.Location, error) {
	if len(r.listLocationsPayload) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsPayload[0]
	r.listLocationsPayload = r.listLocationsPayload[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) UpdateLocationByUid(ctx context.Context, arg db.UpdateLocationByUidParams) (db.Location, error) {
	if len(r.updateLocationByUidPayload) == 0 {
		return db.Location{}, nil
	}

	response := r.updateLocationByUidPayload[0]
	r.updateLocationByUidPayload = r.updateLocationByUidPayload[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) SetGetLocationPayload(response LocationPayload) {
	r.getLocationPayload = append(r.getLocationPayload, response)
}

func (r *MockRepositoryService) SetGetLocationByUidPayload(response LocationPayload) {
	r.getLocationByUidPayload = append(r.getLocationByUidPayload, response)
}

func (r *MockRepositoryService) SetListLocationsPayload(response LocationsPayload) {
	r.listLocationsPayload = append(r.listLocationsPayload, response)
}

func (r *MockRepositoryService) SetUpdateLocationByUidPayload(response LocationPayload) {
	r.updateLocationByUidPayload = append(r.updateLocationByUidPayload, response)
}
