-- name: CreateInvoiceRequest :one
INSERT INTO invoice_requests (
    user_id,
    promotion_id,
    session_id,
    release_date,
    currency,
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
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
  RETURNING *;

-- name: DeleteInvoiceRequest :exec
DELETE FROM invoice_requests
  WHERE id = $1;

-- name: GetInvoiceRequest :one
SELECT * FROM invoice_requests
  WHERE id = $1;

-- name: GetUnsettledInvoiceRequest :one
SELECT * FROM invoice_requests
  WHERE user_id = @user_id::BIGINT AND promotion_id = @promotion_id::BIGINT AND 
    (@memo::TEXT = '' OR @memo::TEXT = memo) AND
    NOT is_settled AND payment_request IS NULL;

-- name: ListUnsettledInvoiceRequests :many
SELECT * FROM invoice_requests
  WHERE user_id = $1 AND is_settled = false AND payment_request IS NULL AND
    (release_date IS NULL OR NOW() > release_date)
  ORDER BY id;

-- name: ListInvoiceRequestsBySessionID :many
SELECT * FROM invoice_requests
  WHERE session_id = $1 AND
    (release_date IS NULL OR NOW() > release_date)
  ORDER BY id;

-- name: ListInvoiceRequestsByPromotionCodeAndSessionID :many
SELECT ir.* FROM invoice_requests ir
  LEFT JOIN promotions p ON p.id = ir.promotion_id
  WHERE ir.session_id = $1 AND p.code = $2 AND
    (ir.release_date IS NULL OR NOW() > ir.release_date)
  ORDER BY ir.id;

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
