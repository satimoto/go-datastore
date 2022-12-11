package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type CommandStopMockData struct {
	CommandStop db.CommandStop
	Error       error
}

type CommandStopsMockData struct {
	CommandStops []db.CommandStop
	Error        error
}

func (r *MockRepositoryService) CreateCommandStop(ctx context.Context, arg db.CreateCommandStopParams) (db.CommandStop, error) {
	r.createCommandStopMockData = append(r.createCommandStopMockData, arg)
	return db.CommandStop{}, nil
}

func (r *MockRepositoryService) GetCommandStop(ctx context.Context, id int64) (db.CommandStop, error) {
	if len(r.getCommandStopMockData) == 0 {
		return db.CommandStop{}, ErrorNotFound()
	}

	response := r.getCommandStopMockData[0]
	r.getCommandStopMockData = r.getCommandStopMockData[1:]
	return response.CommandStop, response.Error
}

func (r *MockRepositoryService) GetCommandStopBySessionID(ctx context.Context, sessionID string) (db.CommandStop, error) {
	if len(r.getCommandStopBySessionIDMockData) == 0 {
		return db.CommandStop{}, ErrorNotFound()
	}

	response := r.getCommandStopBySessionIDMockData[0]
	r.getCommandStopBySessionIDMockData = r.getCommandStopBySessionIDMockData[1:]
	return response.CommandStop, response.Error
}

func (r *MockRepositoryService) UpdateCommandStop(ctx context.Context, arg db.UpdateCommandStopParams) (db.CommandStop, error) {
	r.updateCommandStopMockData = append(r.updateCommandStopMockData, arg)
	return db.CommandStop{}, nil
}

func (r *MockRepositoryService) GetCreateCommandStopMockData() (db.CreateCommandStopParams, error) {
	if len(r.createCommandStopMockData) == 0 {
		return db.CreateCommandStopParams{}, ErrorNotFound()
	}

	response := r.createCommandStopMockData[0]
	r.createCommandStopMockData = r.createCommandStopMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateCommandStopMockData() (db.UpdateCommandStopParams, error) {
	if len(r.updateCommandStopMockData) == 0 {
		return db.UpdateCommandStopParams{}, ErrorNotFound()
	}

	response := r.updateCommandStopMockData[0]
	r.updateCommandStopMockData = r.updateCommandStopMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCommandStopMockData(response CommandStopMockData) {
	r.getCommandStopMockData = append(r.getCommandStopMockData, response)
}

func (r *MockRepositoryService) SetGetCommandStopBySessionIDMockData(response CommandStopMockData) {
	r.getCommandStopBySessionIDMockData = append(r.getCommandStopBySessionIDMockData, response)
}
