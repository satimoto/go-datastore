package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/element"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) element.ElementRepository {
	return element.ElementRepository(repositoryService)
}
