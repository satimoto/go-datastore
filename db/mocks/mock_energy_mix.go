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
	if len(r.createEnergyMixResponse) == 0 {
		return db.EnergyMix{}, ErrorNotFound()
	}

	response := r.createEnergyMixResponse[0]
	r.createEnergyMixResponse = r.createEnergyMixResponse[1:]
	return response.EnergyMix, response.Error
}

func (r *MockRepositoryService) GetEnergyMix(ctx context.Context, id int64) (db.EnergyMix, error) {
	if len(r.getEnergyMixResponse) == 0 {
		return db.EnergyMix{}, ErrorNotFound()
	}

	response := r.getEnergyMixResponse[0]
	r.getEnergyMixResponse = r.getEnergyMixResponse[1:]
	return response.EnergyMix, response.Error
}

func (r *MockRepositoryService) SetCreateEnergyMixResponse(response EnergyMixResponse) {
	r.createEnergyMixResponse = append(r.createEnergyMixResponse, response)
}

func (r *MockRepositoryService) SetGetEnergyMixResponse(response EnergyMixResponse) {
	r.getEnergyMixResponse = append(r.getEnergyMixResponse, response)
}
