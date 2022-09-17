package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type InvoiceRequestMockData struct {
	InvoiceRequest db.InvoiceRequest
	Error          error
}

type InvoiceRequestsMockData struct {
	InvoiceRequests []db.InvoiceRequest
	Error           error
}

func (r *MockRepositoryService) CreateInvoiceRequest(ctx context.Context, arg db.CreateInvoiceRequestParams) (db.InvoiceRequest, error) {
	r.createInvoiceRequestMockData = append(r.createInvoiceRequestMockData, arg)
	return db.InvoiceRequest{}, nil
}

func (r *MockRepositoryService) DeleteInvoiceRequest(ctx context.Context, id int64) error {
	r.deleteInvoiceRequestMockData = append(r.deleteInvoiceRequestMockData, id)
	return nil
}

func (r *MockRepositoryService) GetUnsettledInvoiceRequest(ctx context.Context, arg db.GetUnsettledInvoiceRequestParams) (db.InvoiceRequest, error) {
	if len(r.getUnsettledInvoiceRequestMockData) == 0 {
		return db.InvoiceRequest{}, ErrorNotFound()
	}

	response := r.getUnsettledInvoiceRequestMockData[0]
	r.getUnsettledInvoiceRequestMockData = r.getUnsettledInvoiceRequestMockData[1:]
	return response.InvoiceRequest, response.Error
}

func (r *MockRepositoryService) ListUnsettledInvoiceRequests(ctx context.Context, userID int64) ([]db.InvoiceRequest, error) {
	if len(r.listUnsettledInvoiceRequestsMockData) == 0 {
		return []db.InvoiceRequest{}, nil
	}

	response := r.listUnsettledInvoiceRequestsMockData[0]
	r.listUnsettledInvoiceRequestsMockData = r.listUnsettledInvoiceRequestsMockData[1:]
	return response.InvoiceRequests, response.Error
}

func (r *MockRepositoryService) UpdateInvoiceRequest(ctx context.Context, arg db.UpdateInvoiceRequestParams) (db.InvoiceRequest, error) {
	r.updateInvoiceRequestMockData = append(r.updateInvoiceRequestMockData, arg)
	return db.InvoiceRequest{}, nil
}

func (r *MockRepositoryService) GetCreateInvoiceRequestMockData() (db.CreateInvoiceRequestParams, error) {
	if len(r.createInvoiceRequestMockData) == 0 {
		return db.CreateInvoiceRequestParams{}, ErrorNotFound()
	}

	response := r.createInvoiceRequestMockData[0]
	r.createInvoiceRequestMockData = r.createInvoiceRequestMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteInvoiceRequestMockData() (int64, error) {
	if len(r.deleteInvoiceRequestMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteInvoiceRequestMockData[0]
	r.deleteInvoiceRequestMockData = r.deleteInvoiceRequestMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetUnsettledInvoiceRequestMockData(response InvoiceRequestMockData) {
	r.getUnsettledInvoiceRequestMockData = append(r.getUnsettledInvoiceRequestMockData, response)
}

func (r *MockRepositoryService) SetListInvoiceRequestsMockData(response InvoiceRequestsMockData) {
	r.listUnsettledInvoiceRequestsMockData = append(r.listUnsettledInvoiceRequestsMockData, response)
}

func (r *MockRepositoryService) GetUpdateInvoiceRequestMockData() (db.UpdateInvoiceRequestParams, error) {
	if len(r.updateInvoiceRequestMockData) == 0 {
		return db.UpdateInvoiceRequestParams{}, ErrorNotFound()
	}

	response := r.updateInvoiceRequestMockData[0]
	r.updateInvoiceRequestMockData = r.updateInvoiceRequestMockData[1:]
	return response, nil
}
