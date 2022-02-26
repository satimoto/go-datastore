-- name: CreateChannelRequest :one
INSERT INTO channel_requests (
    user_id,
    status,
    pubkey, 
    payment_hash, 
    payment_addr,
    amount_msat,
    settled_msat
  ) VALUES ($1, $2, $3, $4, $5, $6, $7)
  RETURNING *;

-- name: DeleteChannelRequest :exec
DELETE FROM channel_requests
  WHERE id = $1;

-- name: GetChannelRequest :one
SELECT * FROM channel_requests
  WHERE id = $1;

-- name: GetChannelRequestByPaymentHash :one
SELECT * FROM channel_requests
  WHERE payment_hash = $1 OR sha256('probing-01:' || payment_hash) = $1;

-- name: UpdateChannelRequest :one
UPDATE channel_requests SET (
    status,
    settled_msat,
    funding_tx_id, 
    output_index
  ) = ($2, $3, $4, $5)
  WHERE id = $1
  RETURNING *;

-- name: UpdateChannelRequestStatus :one
UPDATE channel_requests SET status = $2
  WHERE id = $1
  RETURNING *;

-- name: UpdateChannelRequestByChannelPoint :one
UPDATE channel_requests SET status = $3
  WHERE funding_tx_id = $1 AND output_index = $2
  RETURNING *;
