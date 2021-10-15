package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type LocationMockData struct {
	Location db.Location
	Error    error
}

type LocationsMockData struct {
	Locations []db.Location
	Error     error
}

func (r *MockRepositoryService) GetLocation(ctx context.Context, id int64) (db.Location, error) {
	if len(r.getLocationMockData) == 0 {
		return db.Location{}, nil
	}

	response := r.getLocationMockData[0]
	r.getLocationMockData = r.getLocationMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) GetLocationByUid(ctx context.Context, uid string) (db.Location, error) {
	if len(r.getLocationByUidMockData) == 0 {
		return db.Location{}, nil
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

func (r *MockRepositoryService) UpdateLocationByUid(ctx context.Context, arg db.UpdateLocationByUidParams) (db.Location, error) {
	if len(r.updateLocationByUidMockData) == 0 {
		return db.Location{}, nil
	}

	response := r.updateLocationByUidMockData[0]
	r.updateLocationByUidMockData = r.updateLocationByUidMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error {
	if len(r.updateLocationLastUpdatedMockData) == 0 {
		return nil
	}

	response := r.updateLocationLastUpdatedMockData[0]
	r.updateLocationLastUpdatedMockData = r.updateLocationLastUpdatedMockData[1:]
	return response
}

func (r *MockRepositoryService) SetGetLocationMockData(response LocationMockData) {
	r.getLocationMockData = append(r.getLocationMockData, response)
}

func (r *MockRepositoryService) SetGetLocationByUidMockData(response LocationMockData) {
	r.getLocationByUidMockData = append(r.getLocationByUidMockData, response)
}

func (r *MockRepositoryService) SetListLocationsMockData(response LocationsMockData) {
	r.listLocationsMockData = append(r.listLocationsMockData, response)
}

func (r *MockRepositoryService) SetUpdateLocationByUidMockData(response LocationMockData) {
	r.updateLocationByUidMockData = append(r.updateLocationByUidMockData, response)
}

func (r *MockRepositoryService) SetUpdateLocationLastUpdatedMockData(err error) {
	r.updateLocationLastUpdatedMockData = append(r.updateLocationLastUpdatedMockData, err)
}
