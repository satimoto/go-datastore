package mocks

import (
	"github.com/satimoto/go-datastore/pkg/dataimport"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) dataimport.DataImportRepository {
	return dataimport.DataImportRepository(repositoryService)
}
