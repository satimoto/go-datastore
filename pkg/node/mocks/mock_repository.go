package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/node"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) node.NodeRepository {
	return node.NodeRepository(repositoryService)
}
