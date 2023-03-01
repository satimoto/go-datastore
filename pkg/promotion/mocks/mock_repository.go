package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/promotion"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) promotion.PromotionRepository {
	return promotion.PromotionRepository(repositoryService)
}
