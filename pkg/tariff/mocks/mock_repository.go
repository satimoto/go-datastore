package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/tariff"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) tariff.TariffRepository {
	return tariff.TariffRepository(repositoryService)
}
