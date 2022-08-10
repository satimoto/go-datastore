package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/image"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) image.ImageRepository {
	return image.ImageRepository(repositoryService)
}
