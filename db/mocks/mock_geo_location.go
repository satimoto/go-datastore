package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type GeoLocationResponse struct {
	GeoLocation db.GeoLocation
	Error       error
}

type GeoLocationsResponse struct {
	GeoLocations []db.GeoLocation
	Error       error
}

func (r *MockRepositoryService) CreateGeoLocation(ctx context.Context, arg db.CreateGeoLocationParams) (db.GeoLocation, error) {
	return r.createGeoLocationResponse.GeoLocation, r.createGeoLocationResponse.Error
}

func (r *MockRepositoryService) GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error) {
	return r.getGeoLocationResponse.GeoLocation, r.getGeoLocationResponse.Error
}

func (r *MockRepositoryService) SetCreateGeoLocationResponse(response GeoLocationResponse) {
	r.createGeoLocationResponse = response
}

func (r *MockRepositoryService) SetGetGeoLocationResponse(response GeoLocationResponse) {
	r.getGeoLocationResponse = response
}
