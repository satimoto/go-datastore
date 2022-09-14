-- name: CreateInvoiceRequest :one
INSERT INTO invoice_requests (
    user_id,
    promotion_id,
    amount_msat,
    is_settled, 
    payment_request
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: DeleteInvoiceRequest :exec
DELETE FROM invoice_requests
  WHERE id = $1;

-- name: GetUnsettledInvoiceRequestByPromotionCode :one
SELECT ir.* FROM invoice_requests ir
  INNER JOIN promotions p ON p.id = ir.promotion_id
  WHERE p.code = $2 AND ir.user_id = $1 AND NOT ir.is_settled AND ir.payment_request IS NULL;

-- name: ListUnsettledInvoiceRequests :many
SELECT * FROM invoice_requests
  WHERE NOT user_id = $1 AND is_settled AND payment_request IS NULL
  ORDER BY id;

-- name: UpdateInvoiceRequest :one
UPDATE invoice_requests SET (
    amount_msat,
    is_settled,
    payment_request
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;
