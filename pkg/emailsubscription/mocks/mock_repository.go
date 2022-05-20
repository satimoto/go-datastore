package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/emailsubscription"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) emailsubscription.EmailSubscriptionRepository {
	return emailsubscription.EmailSubscriptionRepository(repositoryService)
}
