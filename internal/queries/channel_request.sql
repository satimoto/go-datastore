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
    settled_msat
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
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
    pending_chan_id
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;

-- name: UpdateChannelRequestStatus :one
UPDATE channel_requests SET status = $2
  WHERE id = $1
  RETURNING *;

-- name: UpdatePendingChannelRequestByPubkey :one
UPDATE channel_requests SET status = $3
  WHERE pubkey = $1 AND amount = $2 AND status = 'OPENING_CHANNEL'
  RETURNING *;
