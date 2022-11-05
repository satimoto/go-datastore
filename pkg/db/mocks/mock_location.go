package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type LocationMockData struct {
	Location db.Location
	Error    error
}

type LocationsMockData struct {
	Locations []db.Location
	Error     error
}

func (r *MockRepositoryService) CreateLocation(ctx context.Context, arg db.CreateLocationParams) (db.Location, error) {
	r.createLocationMockData = append(r.createLocationMockData, arg)
	return db.Location{}, nil
}

func (r *MockRepositoryService) GetLocation(ctx context.Context, id int64) (db.Location, error) {
	if len(r.getLocationMockData) == 0 {
		return db.Location{}, ErrorNotFound()
	}

	response := r.getLocationMockData[0]
	r.getLocationMockData = r.getLocationMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) GetLocationByLastUpdated(ctx context.Context, arg db.GetLocationByLastUpdatedParams) (db.Location, error) {
	if len(r.getLocationByLastUpdatedMockData) == 0 {
		return db.Location{}, ErrorNotFound()
	}

	response := r.getLocationByLastUpdatedMockData[0]
	r.getLocationByLastUpdatedMockData = r.getLocationByLastUpdatedMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) GetLocationByUid(ctx context.Context, uid string) (db.Location, error) {
	if len(r.getLocationByUidMockData) == 0 {
		return db.Location{}, ErrorNotFound()
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

func (r *MockRepositoryService) ListLocationsByCountry(ctx context.Context, country string) ([]db.Location, error) {
	if len(r.listLocationsByCountryMockData) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsByCountryMockData[0]
	r.listLocationsByCountryMockData = r.listLocationsByCountryMockData[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) ListLocationsByGeom(ctx context.Context, arg db.ListLocationsByGeomParams) ([]db.Location, error) {
	if len(r.listLocationsByGeomMockData) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsByGeomMockData[0]
	r.listLocationsByGeomMockData = r.listLocationsByGeomMockData[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) UpdateLocationAvailability(ctx context.Context, arg db.UpdateLocationAvailabilityParams) error {
	r.updateLocationAvailabilityMockData = append(r.updateLocationAvailabilityMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationByUid(ctx context.Context, arg db.UpdateLocationByUidParams) (db.Location, error) {
	r.updateLocationByUidMockData = append(r.updateLocationByUidMockData, arg)
	return db.Location{}, nil
}

func (r *MockRepositoryService) UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error {
	r.updateLocationLastUpdatedMockData = append(r.updateLocationLastUpdatedMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationPublish(ctx context.Context, arg db.UpdateLocationPublishParams) error {
	r.updateLocationPublishMockData = append(r.updateLocationPublishMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationsPublishByCredential(ctx context.Context, arg db.UpdateLocationsPublishByCredentialParams) error {
	r.updateLocationsPublishByCredentialMockData = append(r.updateLocationsPublishByCredentialMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationsPublishByPartyAndCountryCode(ctx context.Context, arg db.UpdateLocationsPublishByPartyAndCountryCodeParams) error {
	r.updateLocationsPublishByPartyAndCountryCodeMockData = append(r.updateLocationsPublishByPartyAndCountryCodeMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetCreateLocationMockData() (db.CreateLocationParams, error) {
	if len(r.createLocationMockData) == 0 {
		return db.CreateLocationParams{}, ErrorNotFound()
	}

	response := r.createLocationMockData[0]
	r.createLocationMockData = r.createLocationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationAvailabilityMockData() (db.UpdateLocationAvailabilityParams, error) {
	if len(r.updateLocationAvailabilityMockData) == 0 {
		return db.UpdateLocationAvailabilityParams{}, ErrorNotFound()
	}

	response := r.updateLocationAvailabilityMockData[0]
	r.updateLocationAvailabilityMockData = r.updateLocationAvailabilityMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationByUidMockData() (db.UpdateLocationByUidParams, error) {
	if len(r.updateLocationByUidMockData) == 0 {
		return db.UpdateLocationByUidParams{}, ErrorNotFound()
	}

	response := r.updateLocationByUidMockData[0]
	r.updateLocationByUidMockData = r.updateLocationByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationPublishMockData() (db.UpdateLocationPublishParams, error) {
	if len(r.updateLocationPublishMockData) == 0 {
		return db.UpdateLocationPublishParams{}, ErrorNotFound()
	}

	response := r.updateLocationPublishMockData[0]
	r.updateLocationPublishMockData = r.updateLocationPublishMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationsPublishByCredentialMockData() (db.UpdateLocationsPublishByCredentialParams, error) {
	if len(r.updateLocationsPublishByCredentialMockData) == 0 {
		return db.UpdateLocationsPublishByCredentialParams{}, ErrorNotFound()
	}

	response := r.updateLocationsPublishByCredentialMockData[0]
	r.updateLocationsPublishByCredentialMockData = r.updateLocationsPublishByCredentialMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationsPublishByPartyAndCountryCodeMockData() (db.UpdateLocationsPublishByPartyAndCountryCodeParams, error) {
	if len(r.updateLocationsPublishByPartyAndCountryCodeMockData) == 0 {
		return db.UpdateLocationsPublishByPartyAndCountryCodeParams{}, ErrorNotFound()
	}

	response := r.updateLocationsPublishByPartyAndCountryCodeMockData[0]
	r.updateLocationsPublishByPartyAndCountryCodeMockData = r.updateLocationsPublishByPartyAndCountryCodeMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetLocationMockData(response LocationMockData) {
	r.getLocationMockData = append(r.getLocationMockData, response)
}

func (r *MockRepositoryService) SetGetLocationByLastUpdatedMockData(response LocationMockData) {
	r.getLocationByLastUpdatedMockData = append(r.getLocationByLastUpdatedMockData, response)
}

func (r *MockRepositoryService) SetGetLocationByUidMockData(response LocationMockData) {
	r.getLocationByUidMockData = append(r.getLocationByUidMockData, response)
}

func (r *MockRepositoryService) SetListLocationsMockData(response LocationsMockData) {
	r.listLocationsMockData = append(r.listLocationsMockData, response)
}

func (r *MockRepositoryService) SetListLocationsByCountryMockData(response LocationsMockData) {
	r.listLocationsByCountryMockData = append(r.listLocationsByCountryMockData, response)
}

func (r *MockRepositoryService) SetListLocationsByGeomMockData(response LocationsMockData) {
	r.listLocationsByGeomMockData = append(r.listLocationsByGeomMockData, response)
}
