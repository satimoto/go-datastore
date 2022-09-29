-- name: CreatePendingNotification :one
INSERT INTO pending_notifications (
    user_id,
    node_id,
    invoice_request_id,
    type,
    send_date
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: DeletePendingNotification :exec
DELETE FROM pending_notifications
  WHERE id = $1;

-- name: DeletePendingNotificationByInvoiceRequest :exec
DELETE FROM pending_notifications
  WHERE invoice_request_id = $1
  RETURNING *;

-- name: ListPendingNotifications :many
SELECT * FROM pending_notifications
  WHERE node_id = $1 AND send_date < NOW()
  ORDER BY id;
