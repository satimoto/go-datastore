-- name: CreateChannelRequest :one
INSERT INTO channel_requests (
    user_id,
    node_id,
    status,
    pubkey, 
    payment_hash, 
    payment_addr,
    amount,
    amount_msat,
    settled_msat,
    pending_chan_id
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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

-- name: GetChannelRequestByPendingChanId :one
SELECT * FROM channel_requests
  WHERE pending_chan_id = $1;

-- name: UpdateChannelRequest :one
UPDATE channel_requests SET (
    status,
    settled_msat,
    funding_amount
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;

-- name: UpdateChannelRequestStatus :one
UPDATE channel_requests SET status = $2
  WHERE id = $1
  RETURNING *;

-- name: UpdatePendingChannelRequestByCircuitKey :one
UPDATE channel_requests SET status = $3
  WHERE status = 'OPENING_CHANNEL' AND output_index = $1 AND funding_tx_id = $2
  RETURNING *;

-- name: UpdatePendingChannelRequestByPubkey :one
UPDATE channel_requests SET (
    funding_tx_id, 
    output_index
  ) = ($3, $4)
  WHERE status = 'OPENING_CHANNEL' AND pubkey = $1 AND funding_amount = $2
  RETURNING *;
