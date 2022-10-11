package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type CdrMockData struct {
	Cdr   db.Cdr
	Error error
}

type CdrsMockData struct {
	Cdrs  []db.Cdr
	Error error
}

func (r *MockRepositoryService) CreateCdr(ctx context.Context, arg db.CreateCdrParams) (db.Cdr, error) {
	r.createCdrMockData = append(r.createCdrMockData, arg)
	return db.Cdr{
		Uid:              arg.Uid,
		StartDateTime:    arg.StartDateTime,
		StopDateTime:     arg.StopDateTime,
		AuthID:           arg.AuthID,
		AuthMethod:       arg.AuthMethod,
		MeterID:          arg.MeterID,
		Currency:         arg.Currency,
		TotalCost:        arg.TotalCost,
		TotalEnergy:      arg.TotalEnergy,
		TotalTime:        arg.TotalTime,
		TotalParkingTime: arg.TotalParkingTime,
		Remark:           arg.Remark,
		LastUpdated:      arg.LastUpdated,
	}, nil
}

func (r *MockRepositoryService) GetCdrByLastUpdated(ctx context.Context, arg db.GetCdrByLastUpdatedParams) (db.Cdr, error) {
	if len(r.getCdrByLastUpdatedMockData) == 0 {
		return db.Cdr{}, ErrorNotFound()
	}

	response := r.getCdrByLastUpdatedMockData[0]
	r.getCdrByLastUpdatedMockData = r.getCdrByLastUpdatedMockData[1:]
	return response.Cdr, response.Error
}

func (r *MockRepositoryService) GetCdrByUid(ctx context.Context, uid string) (db.Cdr, error) {
	if len(r.getCdrByUidMockData) == 0 {
		return db.Cdr{}, ErrorNotFound()
	}

	response := r.getCdrByUidMockData[0]
	r.getCdrByUidMockData = r.getCdrByUidMockData[1:]
	return response.Cdr, response.Error
}


func (r *MockRepositoryService) ListCdrsBySessionStatus(ctx context.Context, arg db.ListCdrsBySessionStatusParams) ([]db.Cdr, error) {
	if len(r.listCdrsBySessionStatusMockData) == 0 {
		return []db.Cdr{}, nil
	}

	response := r.listCdrsBySessionStatusMockData[0]
	r.listCdrsBySessionStatusMockData = r.listCdrsBySessionStatusMockData[1:]
	return response.Cdrs, response.Error
}

func (r *MockRepositoryService) GetCreateCdrMockData() (db.CreateCdrParams, error) {
	if len(r.createCdrMockData) == 0 {
		return db.CreateCdrParams{}, ErrorNotFound()
	}

	response := r.createCdrMockData[0]
	r.createCdrMockData = r.createCdrMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCdrByLastUpdatedMockData(response CdrMockData) {
	r.getCdrByLastUpdatedMockData = append(r.getCdrByLastUpdatedMockData, response)
}

func (r *MockRepositoryService) SetGetCdrByUidMockData(response CdrMockData) {
	r.getCdrByUidMockData = append(r.getCdrByUidMockData, response)
}

func (r *MockRepositoryService) SetListCdrsBySessionStatusMockData(response CdrsMockData) {
	r.listCdrsBySessionStatusMockData = append(r.listCdrsBySessionStatusMockData, response)
}
