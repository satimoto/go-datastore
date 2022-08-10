package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/version"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) version.VersionRepository {
	return version.VersionRepository(repositoryService)
}
