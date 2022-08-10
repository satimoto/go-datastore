package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/token"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) token.TokenRepository {
	return token.TokenRepository(repositoryService)
}
