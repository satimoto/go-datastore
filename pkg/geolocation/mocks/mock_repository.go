package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/geolocation"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) geolocation.GeoLocationRepository {
	return geolocation.GeoLocationRepository(repositoryService)
}
