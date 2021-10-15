package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListRelatedLocations(ctx context.Context, locationID int64) ([]db.GeoLocation, error) {
	if len(r.listRelatedLocationsMockData) == 0 {
		return []db.GeoLocation{}, nil
	}

	response := r.listRelatedLocationsMockData[0]
	r.listRelatedLocationsMockData = r.listRelatedLocationsMockData[1:]
	return response.GeoLocations, response.Error
}

func (r *MockRepositoryService) SetListRelatedLocationsMockData(response GeoLocationsMockData) {
	r.listRelatedLocationsMockData = append(r.listRelatedLocationsMockData, response)
}
