package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *MockRepositoryService) DeleteRelatedLocations(ctx context.Context, id int64) error {
	r.deleteRelatedLocationsMockData = append(r.deleteRelatedLocationsMockData, id)
	return nil
}

func (r *MockRepositoryService) ListRelatedLocations(ctx context.Context, locationID int64) ([]db.GeoLocation, error) {
	if len(r.listRelatedLocationsMockData) == 0 {
		return []db.GeoLocation{}, nil
	}

	response := r.listRelatedLocationsMockData[0]
	r.listRelatedLocationsMockData = r.listRelatedLocationsMockData[1:]
	return response.GeoLocations, response.Error
}

func (r *MockRepositoryService) SetRelatedLocation(ctx context.Context, arg db.SetRelatedLocationParams) error {
	r.setRelatedLocationMockData = append(r.setRelatedLocationMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetDeleteRelatedLocationsMockData() (int64, error) {
	if len(r.deleteRelatedLocationsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteRelatedLocationsMockData[0]
	r.deleteRelatedLocationsMockData = r.deleteRelatedLocationsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetSetRelatedLocationMockData() (db.SetRelatedLocationParams, error) {
	if len(r.setRelatedLocationMockData) == 0 {
		return db.SetRelatedLocationParams{}, ErrorNotFound()
	}

	response := r.setRelatedLocationMockData[0]
	r.setRelatedLocationMockData = r.setRelatedLocationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListRelatedLocationsMockData(response GeoLocationsMockData) {
	r.listRelatedLocationsMockData = append(r.listRelatedLocationsMockData, response)
}
