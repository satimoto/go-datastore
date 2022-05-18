package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type LocationMockData struct {
	Location db.Location
	Error    error
}

type LocationsMockData struct {
	Locations []db.Location
	Error     error
}

func (r *MockRepositoryService) CreateLocation(ctx context.Context, arg db.CreateLocationParams) (db.Location, error) {
	r.createLocationMockData = append(r.createLocationMockData, arg)
	return db.Location{}, nil
}

func (r *MockRepositoryService) GetLocation(ctx context.Context, id int64) (db.Location, error) {
	if len(r.getLocationMockData) == 0 {
		return db.Location{}, ErrorNotFound()
	}

	response := r.getLocationMockData[0]
	r.getLocationMockData = r.getLocationMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) GetLocationByLastUpdated(ctx context.Context, arg db.GetLocationByLastUpdatedParams) (db.Location, error) {
	if len(r.GetLocationByLastUpdatedMockData) == 0 {
		return db.Location{}, ErrorNotFound()
	}

	response := r.GetLocationByLastUpdatedMockData[0]
	r.GetLocationByLastUpdatedMockData = r.GetLocationByLastUpdatedMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) GetLocationByUid(ctx context.Context, uid string) (db.Location, error) {
	if len(r.getLocationByUidMockData) == 0 {
		return db.Location{}, ErrorNotFound()
	}

	response := r.getLocationByUidMockData[0]
	r.getLocationByUidMockData = r.getLocationByUidMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) ListLocations(ctx context.Context) ([]db.Location, error) {
	if len(r.listLocationsMockData) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsMockData[0]
	r.listLocationsMockData = r.listLocationsMockData[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) ListLocationsByGeom(ctx context.Context, arg db.ListLocationsByGeomParams) ([]db.Location, error) {
	if len(r.listLocationsByGeomMockData) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsByGeomMockData[0]
	r.listLocationsByGeomMockData = r.listLocationsByGeomMockData[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) UpdateLocationAvailability(ctx context.Context, arg db.UpdateLocationAvailabilityParams) error {
	r.updateLocationAvailabilityMockData = append(r.updateLocationAvailabilityMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationByUid(ctx context.Context, arg db.UpdateLocationByUidParams) (db.Location, error) {
	r.updateLocationByUidMockData = append(r.updateLocationByUidMockData, arg)
	return db.Location{}, nil
}

func (r *MockRepositoryService) UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error {
	r.updateLocationLastUpdatedMockData = append(r.updateLocationLastUpdatedMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetCreateLocationMockData() (db.CreateLocationParams, error) {
	if len(r.createLocationMockData) == 0 {
		return db.CreateLocationParams{}, ErrorNotFound()
	}

	response := r.createLocationMockData[0]
	r.createLocationMockData = r.createLocationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationAvailabilityMockData() (db.UpdateLocationAvailabilityParams, error) {
	if len(r.updateLocationAvailabilityMockData) == 0 {
		return db.UpdateLocationAvailabilityParams{}, ErrorNotFound()
	}

	response := r.updateLocationAvailabilityMockData[0]
	r.updateLocationAvailabilityMockData = r.updateLocationAvailabilityMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationByUidMockData() (db.UpdateLocationByUidParams, error) {
	if len(r.updateLocationByUidMockData) == 0 {
		return db.UpdateLocationByUidParams{}, ErrorNotFound()
	}

	response := r.updateLocationByUidMockData[0]
	r.updateLocationByUidMockData = r.updateLocationByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetLocationMockData(response LocationMockData) {
	r.getLocationMockData = append(r.getLocationMockData, response)
}

func (r *MockRepositoryService) SetGetLocationByLastUpdatedMockData(response LocationMockData) {
	r.GetLocationByLastUpdatedMockData = append(r.GetLocationByLastUpdatedMockData, response)
}

func (r *MockRepositoryService) SetGetLocationByUidMockData(response LocationMockData) {
	r.getLocationByUidMockData = append(r.getLocationByUidMockData, response)
}

func (r *MockRepositoryService) SetListLocationsMockData(response LocationsMockData) {
	r.listLocationsMockData = append(r.listLocationsMockData, response)
}

func (r *MockRepositoryService) SetListLocationsByGeomMockData(response LocationsMockData) {
	r.listLocationsByGeomMockData = append(r.listLocationsByGeomMockData, response)
}
