package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListRelatedLocations(ctx context.Context, locationID int64) ([]db.GeoLocation, error) {
	return r.listRelatedLocationsResponse.GeoLocations, r.listRelatedLocationsResponse.Error
}

func (r *MockRepositoryService) SetListRelatedLocationsResponse(response GeoLocationsResponse) {
	r.listRelatedLocationsResponse = response
}
