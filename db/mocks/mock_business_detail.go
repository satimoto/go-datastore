package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type BusinessDetailResponse struct {
	BusinessDetail db.BusinessDetail
	Error          error
}

func (r *MockRepositoryService) CreateBusinessDetail(ctx context.Context, arg db.CreateBusinessDetailParams) (db.BusinessDetail, error) {
	if len(r.createBusinessDetailResponse) == 0 {
		return db.BusinessDetail{}, ErrorNotFound()
	}
	
	response := r.createBusinessDetailResponse[0]
	r.createBusinessDetailResponse = r.createBusinessDetailResponse[1:]
	return response.BusinessDetail, response.Error
}

func (r *MockRepositoryService) GetBusinessDetail(ctx context.Context, id int64) (db.BusinessDetail, error) {
	if len(r.getBusinessDetailResponse) == 0 {
		return db.BusinessDetail{}, ErrorNotFound()
	}
	
	response := r.getBusinessDetailResponse[0]
	r.getBusinessDetailResponse = r.getBusinessDetailResponse[1:]
	return response.BusinessDetail, response.Error
}

func (r *MockRepositoryService) SetCreateBusinessDetailResponse(response BusinessDetailResponse) {
	r.createBusinessDetailResponse = append(r.createBusinessDetailResponse, response)
}

func (r *MockRepositoryService) SetGetBusinessDetailResponse(response BusinessDetailResponse) {
	r.getBusinessDetailResponse = append(r.getBusinessDetailResponse, response)
}
