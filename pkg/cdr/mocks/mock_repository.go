package mocks

import (
	"github.com/satimoto/go-datastore/pkg/cdr"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) cdr.CdrRepository {
	return cdr.CdrRepository(repositoryService)
}
