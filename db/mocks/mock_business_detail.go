package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type BusinessDetailPayload struct {
	BusinessDetail db.BusinessDetail
	Error          error
}

func (r *MockRepositoryService) CreateBusinessDetail(ctx context.Context, arg db.CreateBusinessDetailParams) (db.BusinessDetail, error) {
	if len(r.createBusinessDetailPayload) == 0 {
		return db.BusinessDetail{}, ErrorNotFound()
	}

	response := r.createBusinessDetailPayload[0]
	r.createBusinessDetailPayload = r.createBusinessDetailPayload[1:]
	return response.BusinessDetail, response.Error
}

func (r *MockRepositoryService) GetBusinessDetail(ctx context.Context, id int64) (db.BusinessDetail, error) {
	if len(r.getBusinessDetailPayload) == 0 {
		return db.BusinessDetail{}, ErrorNotFound()
	}

	response := r.getBusinessDetailPayload[0]
	r.getBusinessDetailPayload = r.getBusinessDetailPayload[1:]
	return response.BusinessDetail, response.Error
}

func (r *MockRepositoryService) SetCreateBusinessDetailPayload(response BusinessDetailPayload) {
	r.createBusinessDetailPayload = append(r.createBusinessDetailPayload, response)
}

func (r *MockRepositoryService) SetGetBusinessDetailPayload(response BusinessDetailPayload) {
	r.getBusinessDetailPayload = append(r.getBusinessDetailPayload, response)
}
