package mocks

import (
	"github.com/satimoto/go-datastore/pkg/connector"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) connector.ConnectorRepository {
	return connector.ConnectorRepository(repositoryService)
}
