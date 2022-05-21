package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/tokenauthorization"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) tokenauthorization.TokenAuthorizationRepository {
	return tokenauthorization.TokenAuthorizationRepository(repositoryService)
}
