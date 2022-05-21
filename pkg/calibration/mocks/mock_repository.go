package mocks

import (
	"github.com/satimoto/go-datastore/pkg/calibration"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) calibration.CalibrationRepository {
	return calibration.CalibrationRepository(repositoryService)
}
