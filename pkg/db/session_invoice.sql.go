// Code generated by sqlc. DO NOT EDIT.
// source: session_invoice.sql

package db

import (
	"context"
	"time"
)

const createSessionInvoice = `-- name: CreateSessionInvoice :one
INSERT INTO session_invoices (
    session_id,
    user_id,
    currency,
    currency_rate,
    currency_rate_msat,
    price_fiat,
    price_msat,
    commission_fiat,
    commission_msat,
    tax_fiat,
    tax_msat,
    total_fiat,
    total_msat,
    payment_request,
    signature,
    is_settled,
    is_expired,
    estimated_energy,
    estimated_time,
    metered_energy,
    metered_time,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
  RETURNING id, session_id, user_id, currency, currency_rate, currency_rate_msat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, payment_request, is_settled, is_expired, last_updated, total_fiat, total_msat, estimated_energy, estimated_time, metered_energy, metered_time, signature
`

type CreateSessionInvoiceParams struct {
	SessionID        int64     `db:"session_id" json:"sessionID"`
	UserID           int64     `db:"user_id" json:"userID"`
	Currency         string    `db:"currency" json:"currency"`
	CurrencyRate     int64     `db:"currency_rate" json:"currencyRate"`
	CurrencyRateMsat int64     `db:"currency_rate_msat" json:"currencyRateMsat"`
	PriceFiat        float64   `db:"price_fiat" json:"priceFiat"`
	PriceMsat        int64     `db:"price_msat" json:"priceMsat"`
	CommissionFiat   float64   `db:"commission_fiat" json:"commissionFiat"`
	CommissionMsat   int64     `db:"commission_msat" json:"commissionMsat"`
	TaxFiat          float64   `db:"tax_fiat" json:"taxFiat"`
	TaxMsat          int64     `db:"tax_msat" json:"taxMsat"`
	TotalFiat        float64   `db:"total_fiat" json:"totalFiat"`
	TotalMsat        int64     `db:"total_msat" json:"totalMsat"`
	PaymentRequest   string    `db:"payment_request" json:"paymentRequest"`
	Signature        string    `db:"signature" json:"signature"`
	IsSettled        bool      `db:"is_settled" json:"isSettled"`
	IsExpired        bool      `db:"is_expired" json:"isExpired"`
	EstimatedEnergy  float64   `db:"estimated_energy" json:"estimatedEnergy"`
	EstimatedTime    float64   `db:"estimated_time" json:"estimatedTime"`
	MeteredEnergy    float64   `db:"metered_energy" json:"meteredEnergy"`
	MeteredTime      float64   `db:"metered_time" json:"meteredTime"`
	LastUpdated      time.Time `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateSessionInvoice(ctx context.Context, arg CreateSessionInvoiceParams) (SessionInvoice, error) {
	row := q.db.QueryRowContext(ctx, createSessionInvoice,
		arg.SessionID,
		arg.UserID,
		arg.Currency,
		arg.CurrencyRate,
		arg.CurrencyRateMsat,
		arg.PriceFiat,
		arg.PriceMsat,
		arg.CommissionFiat,
		arg.CommissionMsat,
		arg.TaxFiat,
		arg.TaxMsat,
		arg.TotalFiat,
		arg.TotalMsat,
		arg.PaymentRequest,
		arg.Signature,
		arg.IsSettled,
		arg.IsExpired,
		arg.EstimatedEnergy,
		arg.EstimatedTime,
		arg.MeteredEnergy,
		arg.MeteredTime,
		arg.LastUpdated,
	)
	var i SessionInvoice
	err := row.Scan(
		&i.ID,
		&i.SessionID,
		&i.UserID,
		&i.Currency,
		&i.CurrencyRate,
		&i.CurrencyRateMsat,
		&i.PriceFiat,
		&i.PriceMsat,
		&i.CommissionFiat,
		&i.CommissionMsat,
		&i.TaxFiat,
		&i.TaxMsat,
		&i.PaymentRequest,
		&i.IsSettled,
		&i.IsExpired,
		&i.LastUpdated,
		&i.TotalFiat,
		&i.TotalMsat,
		&i.EstimatedEnergy,
		&i.EstimatedTime,
		&i.MeteredEnergy,
		&i.MeteredTime,
		&i.Signature,
	)
	return i, err
}

const getSessionInvoice = `-- name: GetSessionInvoice :one
SELECT id, session_id, user_id, currency, currency_rate, currency_rate_msat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, payment_request, is_settled, is_expired, last_updated, total_fiat, total_msat, estimated_energy, estimated_time, metered_energy, metered_time, signature FROM session_invoices
  WHERE id = $1
`

func (q *Queries) GetSessionInvoice(ctx context.Context, id int64) (SessionInvoice, error) {
	row := q.db.QueryRowContext(ctx, getSessionInvoice, id)
	var i SessionInvoice
	err := row.Scan(
		&i.ID,
		&i.SessionID,
		&i.UserID,
		&i.Currency,
		&i.CurrencyRate,
		&i.CurrencyRateMsat,
		&i.PriceFiat,
		&i.PriceMsat,
		&i.CommissionFiat,
		&i.CommissionMsat,
		&i.TaxFiat,
		&i.TaxMsat,
		&i.PaymentRequest,
		&i.IsSettled,
		&i.IsExpired,
		&i.LastUpdated,
		&i.TotalFiat,
		&i.TotalMsat,
		&i.EstimatedEnergy,
		&i.EstimatedTime,
		&i.MeteredEnergy,
		&i.MeteredTime,
		&i.Signature,
	)
	return i, err
}

const getSessionInvoiceByPaymentRequest = `-- name: GetSessionInvoiceByPaymentRequest :one
SELECT id, session_id, user_id, currency, currency_rate, currency_rate_msat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, payment_request, is_settled, is_expired, last_updated, total_fiat, total_msat, estimated_energy, estimated_time, metered_energy, metered_time, signature FROM session_invoices
  WHERE payment_request = $1
`

func (q *Queries) GetSessionInvoiceByPaymentRequest(ctx context.Context, paymentRequest string) (SessionInvoice, error) {
	row := q.db.QueryRowContext(ctx, getSessionInvoiceByPaymentRequest, paymentRequest)
	var i SessionInvoice
	err := row.Scan(
		&i.ID,
		&i.SessionID,
		&i.UserID,
		&i.Currency,
		&i.CurrencyRate,
		&i.CurrencyRateMsat,
		&i.PriceFiat,
		&i.PriceMsat,
		&i.CommissionFiat,
		&i.CommissionMsat,
		&i.TaxFiat,
		&i.TaxMsat,
		&i.PaymentRequest,
		&i.IsSettled,
		&i.IsExpired,
		&i.LastUpdated,
		&i.TotalFiat,
		&i.TotalMsat,
		&i.EstimatedEnergy,
		&i.EstimatedTime,
		&i.MeteredEnergy,
		&i.MeteredTime,
		&i.Signature,
	)
	return i, err
}

const listSessionInvoices = `-- name: ListSessionInvoices :many
SELECT id, session_id, user_id, currency, currency_rate, currency_rate_msat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, payment_request, is_settled, is_expired, last_updated, total_fiat, total_msat, estimated_energy, estimated_time, metered_energy, metered_time, signature FROM session_invoices
  WHERE is_expired = $1 AND is_settled = $2
  ORDER BY id
`

type ListSessionInvoicesParams struct {
	IsExpired bool `db:"is_expired" json:"isExpired"`
	IsSettled bool `db:"is_settled" json:"isSettled"`
}

func (q *Queries) ListSessionInvoices(ctx context.Context, arg ListSessionInvoicesParams) ([]SessionInvoice, error) {
	rows, err := q.db.QueryContext(ctx, listSessionInvoices, arg.IsExpired, arg.IsSettled)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SessionInvoice
	for rows.Next() {
		var i SessionInvoice
		if err := rows.Scan(
			&i.ID,
			&i.SessionID,
			&i.UserID,
			&i.Currency,
			&i.CurrencyRate,
			&i.CurrencyRateMsat,
			&i.PriceFiat,
			&i.PriceMsat,
			&i.CommissionFiat,
			&i.CommissionMsat,
			&i.TaxFiat,
			&i.TaxMsat,
			&i.PaymentRequest,
			&i.IsSettled,
			&i.IsExpired,
			&i.LastUpdated,
			&i.TotalFiat,
			&i.TotalMsat,
			&i.EstimatedEnergy,
			&i.EstimatedTime,
			&i.MeteredEnergy,
			&i.MeteredTime,
			&i.Signature,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSessionInvoicesBySessionID = `-- name: ListSessionInvoicesBySessionID :many
SELECT id, session_id, user_id, currency, currency_rate, currency_rate_msat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, payment_request, is_settled, is_expired, last_updated, total_fiat, total_msat, estimated_energy, estimated_time, metered_energy, metered_time, signature FROM session_invoices
  WHERE session_id = $1
  ORDER BY id
`

func (q *Queries) ListSessionInvoicesBySessionID(ctx context.Context, sessionID int64) ([]SessionInvoice, error) {
	rows, err := q.db.QueryContext(ctx, listSessionInvoicesBySessionID, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SessionInvoice
	for rows.Next() {
		var i SessionInvoice
		if err := rows.Scan(
			&i.ID,
			&i.SessionID,
			&i.UserID,
			&i.Currency,
			&i.CurrencyRate,
			&i.CurrencyRateMsat,
			&i.PriceFiat,
			&i.PriceMsat,
			&i.CommissionFiat,
			&i.CommissionMsat,
			&i.TaxFiat,
			&i.TaxMsat,
			&i.PaymentRequest,
			&i.IsSettled,
			&i.IsExpired,
			&i.LastUpdated,
			&i.TotalFiat,
			&i.TotalMsat,
			&i.EstimatedEnergy,
			&i.EstimatedTime,
			&i.MeteredEnergy,
			&i.MeteredTime,
			&i.Signature,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSessionInvoicesByUserID = `-- name: ListSessionInvoicesByUserID :many
SELECT si.id, si.session_id, si.user_id, si.currency, si.currency_rate, si.currency_rate_msat, si.price_fiat, si.price_msat, si.commission_fiat, si.commission_msat, si.tax_fiat, si.tax_msat, si.payment_request, si.is_settled, si.is_expired, si.last_updated, si.total_fiat, si.total_msat, si.estimated_energy, si.estimated_time, si.metered_energy, si.metered_time, si.signature FROM session_invoices si
  INNER JOIN sessions s ON s.id = si.session_id
  INNER JOIN users u ON u.id = s.user_id
  WHERE u.id = $1 AND si.is_settled = $2 AND si.is_expired = $3
  ORDER BY si.id
`

type ListSessionInvoicesByUserIDParams struct {
	ID        int64 `db:"id" json:"id"`
	IsSettled bool  `db:"is_settled" json:"isSettled"`
	IsExpired bool  `db:"is_expired" json:"isExpired"`
}

func (q *Queries) ListSessionInvoicesByUserID(ctx context.Context, arg ListSessionInvoicesByUserIDParams) ([]SessionInvoice, error) {
	rows, err := q.db.QueryContext(ctx, listSessionInvoicesByUserID, arg.ID, arg.IsSettled, arg.IsExpired)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SessionInvoice
	for rows.Next() {
		var i SessionInvoice
		if err := rows.Scan(
			&i.ID,
			&i.SessionID,
			&i.UserID,
			&i.Currency,
			&i.CurrencyRate,
			&i.CurrencyRateMsat,
			&i.PriceFiat,
			&i.PriceMsat,
			&i.CommissionFiat,
			&i.CommissionMsat,
			&i.TaxFiat,
			&i.TaxMsat,
			&i.PaymentRequest,
			&i.IsSettled,
			&i.IsExpired,
			&i.LastUpdated,
			&i.TotalFiat,
			&i.TotalMsat,
			&i.EstimatedEnergy,
			&i.EstimatedTime,
			&i.MeteredEnergy,
			&i.MeteredTime,
			&i.Signature,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSessionInvoice = `-- name: UpdateSessionInvoice :one
UPDATE session_invoices SET (
    payment_request,
    signature,
    is_settled,
    is_expired,
    last_updated
  ) = ($2, $3, $4, $5, $6)
  WHERE id = $1
  RETURNING id, session_id, user_id, currency, currency_rate, currency_rate_msat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, payment_request, is_settled, is_expired, last_updated, total_fiat, total_msat, estimated_energy, estimated_time, metered_energy, metered_time, signature
`

type UpdateSessionInvoiceParams struct {
	ID             int64     `db:"id" json:"id"`
	PaymentRequest string    `db:"payment_request" json:"paymentRequest"`
	Signature      string    `db:"signature" json:"signature"`
	IsSettled      bool      `db:"is_settled" json:"isSettled"`
	IsExpired      bool      `db:"is_expired" json:"isExpired"`
	LastUpdated    time.Time `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateSessionInvoice(ctx context.Context, arg UpdateSessionInvoiceParams) (SessionInvoice, error) {
	row := q.db.QueryRowContext(ctx, updateSessionInvoice,
		arg.ID,
		arg.PaymentRequest,
		arg.Signature,
		arg.IsSettled,
		arg.IsExpired,
		arg.LastUpdated,
	)
	var i SessionInvoice
	err := row.Scan(
		&i.ID,
		&i.SessionID,
		&i.UserID,
		&i.Currency,
		&i.CurrencyRate,
		&i.CurrencyRateMsat,
		&i.PriceFiat,
		&i.PriceMsat,
		&i.CommissionFiat,
		&i.CommissionMsat,
		&i.TaxFiat,
		&i.TaxMsat,
		&i.PaymentRequest,
		&i.IsSettled,
		&i.IsExpired,
		&i.LastUpdated,
		&i.TotalFiat,
		&i.TotalMsat,
		&i.EstimatedEnergy,
		&i.EstimatedTime,
		&i.MeteredEnergy,
		&i.MeteredTime,
		&i.Signature,
	)
	return i, err
}
