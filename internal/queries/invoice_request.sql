-- name: CreateInvoiceRequest :one
INSERT INTO invoice_requests (
    user_id,
    promotion_id,
    memo,
    price_fiat,
    price_msat,
    commission_fiat,
    commission_msat,
    tax_fiat,
    tax_msat,
    total_fiat,
    total_msat,
    is_settled, 
    payment_request
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
  RETURNING *;

-- name: DeleteInvoiceRequest :exec
DELETE FROM invoice_requests
  WHERE id = $1;


-- name: GetInvoiceRequest :exec
SELECT * FROM invoice_requests
  WHERE id = $1;

-- name: GetUnsettledInvoiceRequest :one
SELECT * FROM invoice_requests
  WHERE user_id = @user_id::BIGINT AND promotion_id = @promotion_id::BIGINT AND 
    (@memo::TEXT = '' OR @memo::TEXT = memo) AND
    NOT is_settled AND payment_request IS NULL;

-- name: ListUnsettledInvoiceRequests :many
SELECT * FROM invoice_requests
  WHERE NOT user_id = $1 AND is_settled AND payment_request IS NULL
  ORDER BY id;

-- name: UpdateInvoiceRequest :one
UPDATE invoice_requests SET (
    price_fiat,
    price_msat,
    commission_fiat,
    commission_msat,
    tax_fiat,
    tax_msat,
    total_fiat,
    total_msat,
    is_settled,
    payment_request
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE id = $1
  RETURNING *;
