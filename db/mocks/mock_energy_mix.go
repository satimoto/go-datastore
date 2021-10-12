package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type EnergyMixPayload struct {
	EnergyMix db.EnergyMix
	Error     error
}

func (r *MockRepositoryService) CreateEnergyMix(ctx context.Context, arg db.CreateEnergyMixParams) (db.EnergyMix, error) {
	if len(r.createEnergyMixPayload) == 0 {
		return db.EnergyMix{}, ErrorNotFound()
	}

	response := r.createEnergyMixPayload[0]
	r.createEnergyMixPayload = r.createEnergyMixPayload[1:]
	return response.EnergyMix, response.Error
}

func (r *MockRepositoryService) GetEnergyMix(ctx context.Context, id int64) (db.EnergyMix, error) {
	if len(r.getEnergyMixPayload) == 0 {
		return db.EnergyMix{}, ErrorNotFound()
	}

	response := r.getEnergyMixPayload[0]
	r.getEnergyMixPayload = r.getEnergyMixPayload[1:]
	return response.EnergyMix, response.Error
}

func (r *MockRepositoryService) SetCreateEnergyMixPayload(response EnergyMixPayload) {
	r.createEnergyMixPayload = append(r.createEnergyMixPayload, response)
}

func (r *MockRepositoryService) SetGetEnergyMixPayload(response EnergyMixPayload) {
	r.getEnergyMixPayload = append(r.getEnergyMixPayload, response)
}
