package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/displaytext"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) displaytext.DisplayTextRepository {
	return displaytext.DisplayTextRepository(repositoryService)
}
