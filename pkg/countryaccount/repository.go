package countryaccount

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type CountryAccountRepository interface {
	GetCountryAccountByCountry(ctx context.Context, country string) (db.CountryAccount, error)
}

func NewRepository(repositoryService *db.RepositoryService) CountryAccountRepository {
	return CountryAccountRepository(repositoryService)
}
