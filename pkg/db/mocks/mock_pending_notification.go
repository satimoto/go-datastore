package mocks

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/pkg/db"
)

type PendingNotificationMockData struct {
	PendingNotification db.PendingNotification
	Error               error
}

type PendingNotificationsMockData struct {
	PendingNotifications []db.PendingNotification
	Error                error
}

func (r *MockRepositoryService) CreatePendingNotification(ctx context.Context, arg db.CreatePendingNotificationParams) (db.PendingNotification, error) {
	r.createPendingNotificationMockData = append(r.createPendingNotificationMockData, arg)
	return db.PendingNotification{}, nil
}

func (r *MockRepositoryService) DeletePendingNotification(ctx context.Context, id int64) error {
	r.deletePendingNotificationMockData = append(r.deletePendingNotificationMockData, id)
	return nil
}

func (r *MockRepositoryService) DeletePendingNotificationByInvoiceRequest(ctx context.Context, invoiceRequestID sql.NullInt64) error {
	r.deletePendingNotificationByInvoiceRequestMockData = append(r.deletePendingNotificationByInvoiceRequestMockData, invoiceRequestID)
	return nil
}

func (r *MockRepositoryService) DeletePendingNotifications(ctx context.Context, ids []int64) error {
	r.deletePendingNotificationsMockData = append(r.deletePendingNotificationsMockData, ids)
	return nil
}

func (r *MockRepositoryService) ListPendingNotifications(ctx context.Context, nodeID int64) ([]db.PendingNotification, error) {
	if len(r.listPendingNotificationsMockData) == 0 {
		return []db.PendingNotification{}, nil
	}

	response := r.listPendingNotificationsMockData[0]
	r.listPendingNotificationsMockData = r.listPendingNotificationsMockData[1:]
	return response.PendingNotifications, response.Error
}

func (r *MockRepositoryService) UpdatePendingNotification(ctx context.Context, arg db.UpdatePendingNotificationParams) error {
	r.updatePendingNotificationMockData = append(r.updatePendingNotificationMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetCreatePendingNotificationMockData() (db.CreatePendingNotificationParams, error) {
	if len(r.createPendingNotificationMockData) == 0 {
		return db.CreatePendingNotificationParams{}, ErrorNotFound()
	}

	response := r.createPendingNotificationMockData[0]
	r.createPendingNotificationMockData = r.createPendingNotificationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeletePendingNotificationMockData() (int64, error) {
	if len(r.deletePendingNotificationMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deletePendingNotificationMockData[0]
	r.deletePendingNotificationMockData = r.deletePendingNotificationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeletePendingNotificationByInvoiceRequestMockData() (sql.NullInt64, error) {
	if len(r.deletePendingNotificationByInvoiceRequestMockData) == 0 {
		return sql.NullInt64{}, ErrorNotFound()
	}

	response := r.deletePendingNotificationByInvoiceRequestMockData[0]
	r.deletePendingNotificationByInvoiceRequestMockData = r.deletePendingNotificationByInvoiceRequestMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeletePendingNotificationsMockData() ([]int64, error) {
	if len(r.deletePendingNotificationsMockData) == 0 {
		return []int64{}, ErrorNotFound()
	}

	response := r.deletePendingNotificationsMockData[0]
	r.deletePendingNotificationsMockData = r.deletePendingNotificationsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdatePendingNotificationMockData() (db.UpdatePendingNotificationParams, error) {
	if len(r.updatePendingNotificationMockData) == 0 {
		return db.UpdatePendingNotificationParams{}, ErrorNotFound()
	}

	response := r.updatePendingNotificationMockData[0]
	r.updatePendingNotificationMockData = r.updatePendingNotificationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListPendingNotificationsMockData(response PendingNotificationsMockData) {
	r.listPendingNotificationsMockData = append(r.listPendingNotificationsMockData, response)
}
