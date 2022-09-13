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
    pending_chan_id,
    scid,
    fee_base_msat,
    fee_proportional_millionths,
    cltv_expiry_delta
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
  RETURNING *;

-- name: DeleteChannelRequest :exec
DELETE FROM channel_requests
  WHERE id = $1;

-- name: GetChannelRequest :one
SELECT * FROM channel_requests
  WHERE id = $1;

-- name: GetChannelRequestByChannelPoint :one
SELECT * FROM channel_requests
  WHERE funding_tx_id_bytes = $1 AND output_index = $2;

-- name: GetChannelRequestByPaymentHash :one
SELECT * FROM channel_requests
  WHERE payment_hash = $1 OR digest('probing-01:' || payment_hash, 'sha256') = $1;

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

-- name: UpdatePendingChannelRequestByPubkey :one
UPDATE channel_requests SET (
    funding_tx_id, 
    funding_tx_id_bytes, 
    output_index
  ) = ($3, $4, $5)
  WHERE status = 'OPENING_CHANNEL' AND pubkey = $1 AND funding_amount = $2
  RETURNING *;
