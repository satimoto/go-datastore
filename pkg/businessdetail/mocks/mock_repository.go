package mocks

import (
	"github.com/satimoto/go-datastore/pkg/businessdetail"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) businessdetail.BusinessDetailRepository {
	return businessdetail.BusinessDetailRepository(repositoryService)
}
