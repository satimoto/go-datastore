package mocks

import (
	"github.com/satimoto/go-datastore/pkg/channelrequest"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) channelrequest.ChannelRequestRepository {
	return channelrequest.ChannelRequestRepository(repositoryService)
}
