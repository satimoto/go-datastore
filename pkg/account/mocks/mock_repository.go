package mocks

import (
	"github.com/satimoto/go-datastore/pkg/account"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) account.AccountRepository {
	return account.AccountRepository(repositoryService)
}
