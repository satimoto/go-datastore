package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/tariffrestriction"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) tariffrestriction.TariffRestrictionRepository {
	return tariffrestriction.TariffRestrictionRepository(repositoryService)
}
