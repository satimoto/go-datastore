package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/evse"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) evse.EvseRepository {
	return evse.EvseRepository(repositoryService)
}
