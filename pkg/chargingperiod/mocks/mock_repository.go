package mocks

import (
	"github.com/satimoto/go-datastore/pkg/chargingperiod"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) chargingperiod.ChargingPeriodRepository {
	return chargingperiod.ChargingPeriodRepository(repositoryService)
}
