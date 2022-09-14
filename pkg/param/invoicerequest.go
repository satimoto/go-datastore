package param

import "github.com/satimoto/go-datastore/pkg/db"

func NewUpdateInvoiceRequestParams(invoiceRequest db.InvoiceRequest) db.UpdateInvoiceRequestParams {
	return db.UpdateInvoiceRequestParams{
		ID:             invoiceRequest.ID,
		AmountMsat:     invoiceRequest.AmountMsat,
		IsSettled:      invoiceRequest.IsSettled,
		PaymentRequest: invoiceRequest.PaymentRequest,
	}
}
