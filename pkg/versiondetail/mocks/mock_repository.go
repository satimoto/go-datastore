package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/versiondetail"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) versiondetail.VersionDetailRepository {
	return versiondetail.VersionDetailRepository(repositoryService)
}
