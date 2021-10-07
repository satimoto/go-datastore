package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type EnergyMixResponse struct {
	EnergyMix db.EnergyMix
	Error     error
}

func (r *MockRepositoryService) CreateEnergyMix(ctx context.Context, arg db.CreateEnergyMixParams) (db.EnergyMix, error) {
	return r.createEnergyMixResponse.EnergyMix, r.createEnergyMixResponse.Error
}

func (r *MockRepositoryService) GetEnergyMix(ctx context.Context, id int64) (db.EnergyMix, error) {
	return r.getEnergyMixResponse.EnergyMix, r.getEnergyMixResponse.Error
}

func (r *MockRepositoryService) SetCreateEnergyMixResponse(response EnergyMixResponse) {
	r.createEnergyMixResponse = response
}

func (r *MockRepositoryService) SetGetEnergyMixResponse(response EnergyMixResponse) {
	r.getEnergyMixResponse = response
}
