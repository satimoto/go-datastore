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
	if len(r.getGetCdrByLastUpdatedMockData) == 0 {
		return db.Cdr{}, ErrorNotFound()
	}

	response := r.getGetCdrByLastUpdatedMockData[0]
	r.getGetCdrByLastUpdatedMockData = r.getGetCdrByLastUpdatedMockData[1:]
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

func (r *MockRepositoryService) GetCreateCdrMockData() (db.CreateCdrParams, error) {
	if len(r.createCdrMockData) == 0 {
		return db.CreateCdrParams{}, ErrorNotFound()
	}

	response := r.createCdrMockData[0]
	r.createCdrMockData = r.createCdrMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCdrByLastUpdatedMockData(response CdrMockData) {
	r.getGetCdrByLastUpdatedMockData = append(r.getGetCdrByLastUpdatedMockData, response)
}

func (r *MockRepositoryService) SetGetCdrByUidMockData(response CdrMockData) {
	r.getCdrByUidMockData = append(r.getCdrByUidMockData, response)
}
