package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListRelatedLocations(ctx context.Context, locationID int64) ([]db.GeoLocation, error) {
	if len(r.listRelatedLocationsPayload) == 0 {
		return []db.GeoLocation{}, nil
	}

	response := r.listRelatedLocationsPayload[0]
	r.listRelatedLocationsPayload = r.listRelatedLocationsPayload[1:]
	return response.GeoLocations, response.Error
}

func (r *MockRepositoryService) SetListRelatedLocationsPayload(response GeoLocationsPayload) {
	r.listRelatedLocationsPayload = append(r.listRelatedLocationsPayload, response)
}
