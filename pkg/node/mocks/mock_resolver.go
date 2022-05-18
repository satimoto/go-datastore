package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/node"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *node.NodeResolver {
	repo := node.NodeRepository(repositoryService)

	return &node.NodeResolver{
		Repository: repo,
	}
}
