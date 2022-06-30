package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type SessionInvoiceMockData struct {
	SessionInvoice db.SessionInvoice
	Error          error
}

type SessionInvoicesMockData struct {
	SessionInvoices []db.SessionInvoice
	Error           error
}

func (r *MockRepositoryService) CreateSessionInvoice(ctx context.Context, arg db.CreateSessionInvoiceParams) (db.SessionInvoice, error) {
	r.createSessionInvoiceMockData = append(r.createSessionInvoiceMockData, arg)
	return db.SessionInvoice{}, nil
}

func (r *MockRepositoryService) GetSessionInvoice(ctx context.Context, id int64) (db.SessionInvoice, error) {
	if len(r.getSessionInvoiceMockData) == 0 {
		return db.SessionInvoice{}, ErrorNotFound()
	}

	response := r.getSessionInvoiceMockData[0]
	r.getSessionInvoiceMockData = r.getSessionInvoiceMockData[1:]
	return response.SessionInvoice, response.Error
}

func (r *MockRepositoryService) GetSessionInvoiceByPaymentRequest(ctx context.Context, paymentRequest string) (db.SessionInvoice, error) {
	if len(r.getSessionInvoiceByPaymentRequestMockData) == 0 {
		return db.SessionInvoice{}, ErrorNotFound()
	}

	response := r.getSessionInvoiceByPaymentRequestMockData[0]
	r.getSessionInvoiceByPaymentRequestMockData = r.getSessionInvoiceByPaymentRequestMockData[1:]
	return response.SessionInvoice, response.Error
}

func (r *MockRepositoryService) ListSessionInvoices(ctx context.Context, sessionID int64) ([]db.SessionInvoice, error) {
	if len(r.listSessionInvoicesMockData) == 0 {
		return []db.SessionInvoice{}, nil
	}

	response := r.listSessionInvoicesMockData[0]
	r.listSessionInvoicesMockData = r.listSessionInvoicesMockData[1:]
	return response.SessionInvoices, response.Error
}

func (r *MockRepositoryService) ListUnsettledSessionInvoicesByUserID(ctx context.Context, id int64) ([]db.SessionInvoice, error) {
	if len(r.listUnsettledSessionInvoicesByUserIDMockData) == 0 {
		return []db.SessionInvoice{}, nil
	}

	response := r.listUnsettledSessionInvoicesByUserIDMockData[0]
	r.listUnsettledSessionInvoicesByUserIDMockData = r.listUnsettledSessionInvoicesByUserIDMockData[1:]
	return response.SessionInvoices, response.Error
}

func (r *MockRepositoryService) UpdateSessionInvoice(ctx context.Context, arg db.UpdateSessionInvoiceParams) (db.SessionInvoice, error) {
	r.updateSessionInvoiceMockData = append(r.updateSessionInvoiceMockData, arg)
	return db.SessionInvoice{}, nil
}

func (r *MockRepositoryService) GetCreateSessionInvoiceMockData() (db.CreateSessionInvoiceParams, error) {
	if len(r.createSessionInvoiceMockData) == 0 {
		return db.CreateSessionInvoiceParams{}, ErrorNotFound()
	}

	response := r.createSessionInvoiceMockData[0]
	r.createSessionInvoiceMockData = r.createSessionInvoiceMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateSessionInvoiceMockData() (db.UpdateSessionInvoiceParams, error) {
	if len(r.updateSessionInvoiceMockData) == 0 {
		return db.UpdateSessionInvoiceParams{}, ErrorNotFound()
	}

	response := r.updateSessionInvoiceMockData[0]
	r.updateSessionInvoiceMockData = r.updateSessionInvoiceMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetSessionInvoiceMockData(response SessionInvoiceMockData) {
	r.getSessionInvoiceMockData = append(r.getSessionInvoiceMockData, response)
}

func (r *MockRepositoryService) SetGetSessionInvoiceByPaymentRequestMockData(response SessionInvoiceMockData) {
	r.getSessionInvoiceByPaymentRequestMockData = append(r.getSessionInvoiceByPaymentRequestMockData, response)
}

func (r *MockRepositoryService) SetListSessionInvoicesMockData(response SessionInvoicesMockData) {
	r.listSessionInvoicesMockData = append(r.listSessionInvoicesMockData, response)
}

func (r *MockRepositoryService) SetListUnsettledSessionInvoicesByUserIDMockData(response SessionInvoicesMockData) {
	r.listUnsettledSessionInvoicesByUserIDMockData = append(r.listUnsettledSessionInvoicesByUserIDMockData, response)
}
