package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type GeoLocationPayload struct {
	GeoLocation db.GeoLocation
	Error       error
}

type GeoLocationsPayload struct {
	GeoLocations []db.GeoLocation
	Error        error
}

func (r *MockRepositoryService) CreateGeoLocation(ctx context.Context, arg db.CreateGeoLocationParams) (db.GeoLocation, error) {
	if len(r.createGeoLocationPayload) == 0 {
		return db.GeoLocation{}, ErrorNotFound()
	}

	response := r.createGeoLocationPayload[0]
	r.createGeoLocationPayload = r.createGeoLocationPayload[1:]
	return response.GeoLocation, response.Error
}

func (r *MockRepositoryService) GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error) {
	if len(r.getGeoLocationPayload) == 0 {
		return db.GeoLocation{}, ErrorNotFound()
	}

	response := r.getGeoLocationPayload[0]
	r.getGeoLocationPayload = r.getGeoLocationPayload[1:]
	return response.GeoLocation, response.Error
}

func (r *MockRepositoryService) SetCreateGeoLocationPayload(response GeoLocationPayload) {
	r.createGeoLocationPayload = append(r.createGeoLocationPayload, response)
}

func (r *MockRepositoryService) SetGetGeoLocationPayload(response GeoLocationPayload) {
	r.getGeoLocationPayload = append(r.getGeoLocationPayload, response)
}
