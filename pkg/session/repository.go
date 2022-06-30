package session

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error)
	CreateSessionInvoice(ctx context.Context, arg db.CreateSessionInvoiceParams) (db.SessionInvoice, error)
	DeleteSessionChargingPeriods(ctx context.Context, sessionID int64) error
	GetSessionByAuthorizationID(ctx context.Context, authorizationID string) (db.Session, error)
	GetSessionByLastUpdated(ctx context.Context, arg db.GetSessionByLastUpdatedParams) (db.Session, error)
	GetSessionByUid(ctx context.Context, uid string) (db.Session, error)
	GetSessionInvoice(ctx context.Context, id int64) (db.SessionInvoice, error)
	GetSessionInvoiceByPaymentRequest(ctx context.Context, paymentRequest string) (db.SessionInvoice, error)
	GetUserBySessionID(ctx context.Context, id int64) (db.User, error)
	ListChargingPeriodDimensions(ctx context.Context, chargingPeriodID int64) ([]db.ChargingPeriodDimension, error)
	ListSessionChargingPeriods(ctx context.Context, sessionID int64) ([]db.ChargingPeriod, error)
	ListSessionInvoices(ctx context.Context, sessionID int64) ([]db.SessionInvoice, error)
	ListUnsettledSessionInvoicesByUserID(ctx context.Context, userID int64) ([]db.SessionInvoice, error)
	SetSessionChargingPeriod(ctx context.Context, arg db.SetSessionChargingPeriodParams) error
	UpdateSessionByUid(ctx context.Context, arg db.UpdateSessionByUidParams) (db.Session, error)
	UpdateSessionInvoice(ctx context.Context, arg db.UpdateSessionInvoiceParams) (db.SessionInvoice, error)
}

func NewRepository(repositoryService *db.RepositoryService) SessionRepository {
	return SessionRepository(repositoryService)
}
