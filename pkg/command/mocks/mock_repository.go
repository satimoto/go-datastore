package mocks

import (
	"github.com/satimoto/go-datastore/pkg/command"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) command.CommandRepository {
	return command.CommandRepository(repositoryService)
}
