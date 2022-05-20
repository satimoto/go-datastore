package mocks

import (
	"github.com/satimoto/go-datastore/pkg/authentication"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) authentication.AuthenticationRepository {
	return authentication.AuthenticationRepository(repositoryService)
}
