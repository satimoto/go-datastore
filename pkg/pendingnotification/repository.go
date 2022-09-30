package pendingnotification

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/pkg/db"
)

type PendingNotificationRepository interface {
	CreatePendingNotification(ctx context.Context, arg db.CreatePendingNotificationParams) (db.PendingNotification, error)
	DeletePendingNotification(ctx context.Context, id int64) error
	DeletePendingNotificationByInvoiceRequest(ctx context.Context, invoiceRequestID sql.NullInt64) error
	DeletePendingNotifications(ctx context.Context, ids []int64) error
	ListPendingNotifications(ctx context.Context, nodeID int64) ([]db.PendingNotification, error)
	UpdatePendingNotification(ctx context.Context, arg db.UpdatePendingNotificationParams) error 
}

func NewRepository(repositoryService *db.RepositoryService) PendingNotificationRepository {
	return PendingNotificationRepository(repositoryService)
}
