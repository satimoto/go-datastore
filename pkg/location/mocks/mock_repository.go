package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/location"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) location.LocationRepository {
	return location.LocationRepository(repositoryService)
}
