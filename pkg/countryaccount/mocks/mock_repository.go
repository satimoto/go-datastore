package mocks

import (
	"github.com/satimoto/go-datastore/pkg/countryaccount"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) countryaccount.CountryAccountRepository {
	return countryaccount.CountryAccountRepository(repositoryService)
}
