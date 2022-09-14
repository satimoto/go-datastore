package invoicerequest

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type InvoiceRequestRepository interface {
	CreateInvoiceRequest(ctx context.Context, arg db.CreateInvoiceRequestParams) (db.InvoiceRequest, error)
	DeleteInvoiceRequest(ctx context.Context, id int64) error
	GetUnsettledInvoiceRequestByPromotionCode(ctx context.Context, arg db.GetUnsettledInvoiceRequestByPromotionCodeParams) (db.InvoiceRequest, error)
	ListUnsettledInvoiceRequests(ctx context.Context, userID int64) ([]db.InvoiceRequest, error)
	UpdateInvoiceRequest(ctx context.Context, arg db.UpdateInvoiceRequestParams) (db.InvoiceRequest, error)
}

func NewRepository(repositoryService *db.RepositoryService) InvoiceRequestRepository {
	return InvoiceRequestRepository(repositoryService)
}
