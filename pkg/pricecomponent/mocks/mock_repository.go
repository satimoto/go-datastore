package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/pricecomponent"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) pricecomponent.PriceComponentRepository {
	return pricecomponent.PriceComponentRepository(repositoryService)
}
