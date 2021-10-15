package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type BusinessDetailMockData struct {
	BusinessDetail db.BusinessDetail
	Error          error
}

func (r *MockRepositoryService) CreateBusinessDetail(ctx context.Context, arg db.CreateBusinessDetailParams) (db.BusinessDetail, error) {
	r.createBusinessDetailMockData = append(r.createBusinessDetailMockData, arg)
	return db.BusinessDetail{}, nil
}

func (r *MockRepositoryService) GetBusinessDetail(ctx context.Context, id int64) (db.BusinessDetail, error) {
	if len(r.getBusinessDetailMockData) == 0 {
		return db.BusinessDetail{}, ErrorNotFound()
	}

	response := r.getBusinessDetailMockData[0]
	r.getBusinessDetailMockData = r.getBusinessDetailMockData[1:]
	return response.BusinessDetail, response.Error
}

func (r *MockRepositoryService) GetCreateBusinessDetailMockData() (db.CreateBusinessDetailParams, error) {
	if len(r.createBusinessDetailMockData) == 0 {
		return db.CreateBusinessDetailParams{}, ErrorNotFound()
	}

	response := r.createBusinessDetailMockData[0]
	r.createBusinessDetailMockData = r.createBusinessDetailMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetBusinessDetailMockData(response BusinessDetailMockData) {
	r.getBusinessDetailMockData = append(r.getBusinessDetailMockData, response)
}
