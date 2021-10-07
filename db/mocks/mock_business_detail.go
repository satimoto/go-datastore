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
	return r.createBusinessDetailResponse.BusinessDetail, r.createBusinessDetailResponse.Error
}

func (r *MockRepositoryService) GetBusinessDetail(ctx context.Context, id int64) (db.BusinessDetail, error) {
	return r.getBusinessDetailResponse.BusinessDetail, r.getBusinessDetailResponse.Error
}

func (r *MockRepositoryService) SetCreateBusinessDetailResponse(response BusinessDetailResponse) {
	r.createBusinessDetailResponse = response
}

func (r *MockRepositoryService) SetGetBusinessDetailResponse(response BusinessDetailResponse) {
	r.getBusinessDetailResponse = response
}
