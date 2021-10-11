package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListRelatedLocations(ctx context.Context, locationID int64) ([]db.GeoLocation, error) {
	if len(r.listRelatedLocationsResponse) == 0 {
		return []db.GeoLocation{}, nil
	}

	response := r.listRelatedLocationsResponse[0]
	r.listRelatedLocationsResponse = r.listRelatedLocationsResponse[1:]
	return response.GeoLocations, response.Error
}

func (r *MockRepositoryService) SetListRelatedLocationsResponse(response GeoLocationsResponse) {
	r.listRelatedLocationsResponse = append(r.listRelatedLocationsResponse, response)
}
