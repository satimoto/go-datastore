-- name: CreateChannelRequest :one
INSERT INTO channel_requests (
    pubkey, 
    payment_hash, 
    payment_addr,
    amount_msat
  ) VALUES ($1, $2, $3, $4)
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
    funding_tx_id, 
    output_index
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
