package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/user"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) user.UserRepository {
	return user.UserRepository(repositoryService)
}
