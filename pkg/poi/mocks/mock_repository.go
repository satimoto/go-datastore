package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/poi"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) poi.PoiRepository {
	return poi.PoiRepository(repositoryService)
}
