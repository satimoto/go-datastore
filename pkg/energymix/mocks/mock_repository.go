package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/energymix"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) energymix.EnergyMixRepository {
	return energymix.EnergyMixRepository(repositoryService)
}
