-- name: CreateChannelRequestHtlc :one
INSERT INTO channel_request_htlcs (
    channel_request_id,
    chan_id, 
    htlc_id, 
    amount_msat,
    is_settled,
    is_failed
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: GetChannelRequestHtlc :one
SELECT * FROM channel_request_htlcs
  WHERE channel_request_id = $1;

-- name: GetChannelRequestHtlcByCircuitKey :one
SELECT * FROM channel_request_htlcs
  WHERE chan_id = $1 AND htlc_id = $2;

-- name: ListChannelRequestHtlcs :many
SELECT * FROM channel_request_htlcs
  WHERE channel_request_id = $1
  ORDER BY id;

-- name: UpdateChannelRequestHtlcByCircuitKey :one
UPDATE channel_request_htlcs SET (
    is_settled, 
    is_failed
  ) = ($3, $4)
  WHERE chan_id = $1 AND htlc_id = $2
  RETURNING *;
