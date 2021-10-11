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
	if len(r.createGeoLocationResponse) == 0 {
		return db.GeoLocation{}, ErrorNotFound()
	}

	response := r.createGeoLocationResponse[0]
	r.createGeoLocationResponse = r.createGeoLocationResponse[1:]
	return response.GeoLocation, response.Error
}

func (r *MockRepositoryService) GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error) {
	if len(r.getGeoLocationResponse) == 0 {
		
		return db.GeoLocation{}, ErrorNotFound()
	}
	response := r.getGeoLocationResponse[0]
	r.getGeoLocationResponse = r.getGeoLocationResponse[1:]
	return response.GeoLocation, response.Error
}

func (r *MockRepositoryService) SetCreateGeoLocationResponse(response GeoLocationResponse) {
	r.createGeoLocationResponse = append(r.createGeoLocationResponse, response)
}

func (r *MockRepositoryService) SetGetGeoLocationResponse(response GeoLocationResponse) {
	r.getGeoLocationResponse = append(r.getGeoLocationResponse, response)
}
