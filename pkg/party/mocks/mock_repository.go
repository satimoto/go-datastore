package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/party"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) party.PartyRepository {
	return party.PartyRepository(repositoryService)
}
