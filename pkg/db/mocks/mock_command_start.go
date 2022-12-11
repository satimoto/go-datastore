package mocks

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/pkg/db"
)

type CommandStartMockData struct {
	CommandStart db.CommandStart
	Error        error
}

type CommandStartsMockData struct {
	CommandStarts []db.CommandStart
	Error         error
}

func (r *MockRepositoryService) CreateCommandStart(ctx context.Context, arg db.CreateCommandStartParams) (db.CommandStart, error) {
	r.createCommandStartMockData = append(r.createCommandStartMockData, arg)
	return db.CommandStart{}, nil
}

func (r *MockRepositoryService) GetCommandStart(ctx context.Context, id int64) (db.CommandStart, error) {
	if len(r.getCommandStartMockData) == 0 {
		return db.CommandStart{}, ErrorNotFound()
	}

	response := r.getCommandStartMockData[0]
	r.getCommandStartMockData = r.getCommandStartMockData[1:]
	return response.CommandStart, response.Error
}

func (r *MockRepositoryService) GetCommandStartByAuthorizationID(ctx context.Context, authorizationID sql.NullString) (db.CommandStart, error) {
	if len(r.getCommandStartByAuthorizationIDMockData) == 0 {
		return db.CommandStart{}, ErrorNotFound()
	}

	response := r.getCommandStartByAuthorizationIDMockData[0]
	r.getCommandStartByAuthorizationIDMockData = r.getCommandStartByAuthorizationIDMockData[1:]
	return response.CommandStart, response.Error
}

func (r *MockRepositoryService) UpdateCommandStart(ctx context.Context, arg db.UpdateCommandStartParams) (db.CommandStart, error) {
	r.updateCommandStartMockData = append(r.updateCommandStartMockData, arg)
	return db.CommandStart{}, nil
}

func (r *MockRepositoryService) GetCreateCommandStartMockData() (db.CreateCommandStartParams, error) {
	if len(r.createCommandStartMockData) == 0 {
		return db.CreateCommandStartParams{}, ErrorNotFound()
	}

	response := r.createCommandStartMockData[0]
	r.createCommandStartMockData = r.createCommandStartMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateCommandStartMockData() (db.UpdateCommandStartParams, error) {
	if len(r.updateCommandStartMockData) == 0 {
		return db.UpdateCommandStartParams{}, ErrorNotFound()
	}

	response := r.updateCommandStartMockData[0]
	r.updateCommandStartMockData = r.updateCommandStartMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCommandStartMockData(response CommandStartMockData) {
	r.getCommandStartMockData = append(r.getCommandStartMockData, response)
}

func (r *MockRepositoryService) SetGetCommandStartByAuthorizationIDMockData(response CommandStartMockData) {
	r.getCommandStartByAuthorizationIDMockData = append(r.getCommandStartByAuthorizationIDMockData, response)
}
