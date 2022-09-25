package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/referral"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) referral.ReferralRepository {
	return referral.ReferralRepository(repositoryService)
}
