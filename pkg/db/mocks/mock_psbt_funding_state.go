package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type PsbtFundingStateMockData struct {
	PsbtFundingState db.PsbtFundingState
	Error            error
}

type PsbtFundingStatesMockData struct {
	PsbtFundingStates []db.PsbtFundingState
	Error             error
}

func (r *MockRepositoryService) CreatePsbtFundingState(ctx context.Context, arg db.CreatePsbtFundingStateParams) (db.PsbtFundingState, error) {
	r.createPsbtFundingStateMockData = append(r.createPsbtFundingStateMockData, arg)
	return db.PsbtFundingState{}, nil
}

func (r *MockRepositoryService) ListPsbtFundingState(ctx context.Context) ([]db.PsbtFundingState, error) {
	if len(r.listPsbtFundingStatesMockData) == 0 {
		return []db.PsbtFundingState{}, nil
	}

	response := r.listPsbtFundingStatesMockData[0]
	r.listPsbtFundingStatesMockData = r.listPsbtFundingStatesMockData[1:]
	return response.PsbtFundingStates, response.Error
}

func (r *MockRepositoryService) UpdatePsbtFundingState(ctx context.Context, arg db.UpdatePsbtFundingStateParams) (db.PsbtFundingState, error) {
	r.updatePsbtFundingStateMockData = append(r.updatePsbtFundingStateMockData, arg)
	return db.PsbtFundingState{}, nil
}

func (r *MockRepositoryService) GetCreatePsbtFundingStateMockData() (db.CreatePsbtFundingStateParams, error) {
	if len(r.createPsbtFundingStateMockData) == 0 {
		return db.CreatePsbtFundingStateParams{}, ErrorNotFound()
	}

	response := r.createPsbtFundingStateMockData[0]
	r.createPsbtFundingStateMockData = r.createPsbtFundingStateMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdatePsbtFundingStateMockData() (db.UpdatePsbtFundingStateParams, error) {
	if len(r.updatePsbtFundingStateMockData) == 0 {
		return db.UpdatePsbtFundingStateParams{}, ErrorNotFound()
	}

	response := r.updatePsbtFundingStateMockData[0]
	r.updatePsbtFundingStateMockData = r.updatePsbtFundingStateMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListPsbtFundingStatesMockData(response PsbtFundingStatesMockData) {
	r.listPsbtFundingStatesMockData = append(r.listPsbtFundingStatesMockData, response)
}
