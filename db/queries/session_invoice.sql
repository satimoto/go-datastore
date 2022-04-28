-- name: CreateSessionInvoice :one
INSERT INTO session_invoices (
    session_id,
    amount_fiat,
    amount_msat,
    commission_fiat,
    commission_msat,
    tax_fiat,
    tax_msat,
    currency,
    payment_request,
    settled,
    expired,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
  RETURNING *;

-- name: GetSessionInvoiceByPaymentRequest :one
SELECT * FROM session_invoices
  WHERE payment_request = $1;

-- name: ListSessionInvoices :many
SELECT * FROM session_invoices
  WHERE session_id = $1
  ORDER BY id;

-- name: UpdateSessionInvoice :one
UPDATE session_invoices SET (
    settled,
    expired,
    last_updated
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;
