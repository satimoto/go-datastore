package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/openingtime"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) openingtime.OpeningTimeRepository {
	return openingtime.OpeningTimeRepository(repositoryService)
}
