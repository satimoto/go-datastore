package mocks

import (
	"github.com/satimoto/go-datastore/pkg/businessdetail"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) businessdetail.BusinessDetailRepository {
	return businessdetail.BusinessDetailRepository(repositoryService)
}
