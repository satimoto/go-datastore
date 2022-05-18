package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type EvseMockData struct {
	Evse  db.Evse
	Error error
}

type EvsesMockData struct {
	Evses []db.Evse
	Error error
}

func (r *MockRepositoryService) CreateEvse(ctx context.Context, arg db.CreateEvseParams) (db.Evse, error) {
	r.createEvseMockData = append(r.createEvseMockData, arg)
	return db.Evse{}, nil
}

func (r *MockRepositoryService) GetEvse(ctx context.Context, id int64) (db.Evse, error) {
	if len(r.getEvseMockData) == 0 {
		return db.Evse{}, ErrorNotFound()
	}

	response := r.getEvseMockData[0]
	r.getEvseMockData = r.getEvseMockData[1:]
	return response.Evse, response.Error
}

func (r *MockRepositoryService) GetEvseByUid(ctx context.Context, uid string) (db.Evse, error) {
	if len(r.getEvseByUidMockData) == 0 {
		return db.Evse{}, ErrorNotFound()
	}

	response := r.getEvseByUidMockData[0]
	r.getEvseByUidMockData = r.getEvseByUidMockData[1:]
	return response.Evse, response.Error
}

func (r *MockRepositoryService) ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error) {
	if len(r.listEvsesMockData) == 0 {
		return []db.Evse{}, nil
	}

	response := r.listEvsesMockData[0]
	r.listEvsesMockData = r.listEvsesMockData[1:]
	return response.Evses, response.Error
}

func (r *MockRepositoryService) UpdateEvseByUid(ctx context.Context, arg db.UpdateEvseByUidParams) (db.Evse, error) {
	r.updateEvseByUidMockData = append(r.updateEvseByUidMockData, arg)
	return db.Evse{}, nil
}

func (r *MockRepositoryService) UpdateEvseLastUpdated(ctx context.Context, arg db.UpdateEvseLastUpdatedParams) error {
	r.updateEvseLastUpdatedMockData = append(r.updateEvseLastUpdatedMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetCreateEvseMockData() (db.CreateEvseParams, error) {
	if len(r.createEvseMockData) == 0 {
		return db.CreateEvseParams{}, ErrorNotFound()
	}

	response := r.createEvseMockData[0]
	r.createEvseMockData = r.createEvseMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateEvseByUidMockData() (db.UpdateEvseByUidParams, error) {
	if len(r.updateEvseByUidMockData) == 0 {
		return db.UpdateEvseByUidParams{}, ErrorNotFound()
	}

	response := r.updateEvseByUidMockData[0]
	r.updateEvseByUidMockData = r.updateEvseByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetUpdateEvseLastUpdatedMockData() (db.UpdateEvseLastUpdatedParams, error) {
	if len(r.updateEvseLastUpdatedMockData) == 0 {
		return db.UpdateEvseLastUpdatedParams{}, ErrorNotFound()
	}

	response := r.updateEvseLastUpdatedMockData[0]
	r.updateEvseLastUpdatedMockData = r.updateEvseLastUpdatedMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetEvseMockData(response EvseMockData) {
	r.getEvseMockData = append(r.getEvseMockData, response)
}

func (r *MockRepositoryService) SetGetEvseByUidMockData(response EvseMockData) {
	r.getEvseByUidMockData = append(r.getEvseByUidMockData, response)
}

func (r *MockRepositoryService) SetListEvsesMockData(response EvsesMockData) {
	r.listEvsesMockData = append(r.listEvsesMockData, response)
}
