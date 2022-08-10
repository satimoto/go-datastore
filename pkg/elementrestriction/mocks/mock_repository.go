package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/elementrestriction"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) elementrestriction.ElementRestrictionRepository {
	return elementrestriction.ElementRestrictionRepository(repositoryService)
}
