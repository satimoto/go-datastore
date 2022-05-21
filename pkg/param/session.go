package param

import (
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
)

func NewUpdateSessionByUidParams(session db.Session) db.UpdateSessionByUidParams {
	return db.UpdateSessionByUidParams{
		Uid:             session.Uid,
		AuthorizationID: session.AuthorizationID,
		StartDatetime:   session.StartDatetime,
		EndDatetime:     session.EndDatetime,
		Kwh:             session.Kwh,
		AuthMethod:      session.AuthMethod,
		MeterID:         session.MeterID,
		Currency:        session.Currency,
		TotalCost:       session.TotalCost,
		Status:          session.Status,
		LastUpdated:     session.LastUpdated,
	}
}

func NewCreateSessionInvoiceParams(session db.Session) db.CreateSessionInvoiceParams {
	return db.CreateSessionInvoiceParams{
		SessionID:   session.ID,
		Currency:    session.Currency,
		IsSettled:   false,
		IsExpired:   false,
		LastUpdated: time.Now(),
	}
}

func NewUpdateSessionInvoiceParams(sessionInvoice db.SessionInvoice) db.UpdateSessionInvoiceParams {
	return db.UpdateSessionInvoiceParams{
		ID:          sessionInvoice.ID,
		IsSettled:   sessionInvoice.IsSettled,
		IsExpired:   sessionInvoice.IsExpired,
		LastUpdated: time.Now(),
	}
}
