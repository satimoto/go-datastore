package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type GeoLocationMockData struct {
	GeoLocation db.GeoLocation
	Error       error
}

type GeoLocationsMockData struct {
	GeoLocations []db.GeoLocation
	Error        error
}

func (r *MockRepositoryService) CreateGeoLocation(ctx context.Context, arg db.CreateGeoLocationParams) (db.GeoLocation, error) {
	r.createGeoLocationMockData = append(r.createGeoLocationMockData, arg)
	return db.GeoLocation{}, nil
}

func (r *MockRepositoryService) GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error) {
	if len(r.getGeoLocationMockData) == 0 {
		return db.GeoLocation{}, ErrorNotFound()
	}

	response := r.getGeoLocationMockData[0]
	r.getGeoLocationMockData = r.getGeoLocationMockData[1:]
	return response.GeoLocation, response.Error
}

func (r *MockRepositoryService) GetCreateGeoLocationMockData() (db.CreateGeoLocationParams, error) {
	if len(r.createGeoLocationMockData) == 0 {
		return db.CreateGeoLocationParams{}, ErrorNotFound()
	}

	response := r.createGeoLocationMockData[0]
	r.createGeoLocationMockData = r.createGeoLocationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetGeoLocationMockData(response GeoLocationMockData) {
	r.getGeoLocationMockData = append(r.getGeoLocationMockData, response)
}
