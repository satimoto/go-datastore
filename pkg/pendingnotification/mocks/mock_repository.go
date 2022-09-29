package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/pendingnotification"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) pendingnotification.PendingNotificationRepository {
	return pendingnotification.PendingNotificationRepository(repositoryService)
}
