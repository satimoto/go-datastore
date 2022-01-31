-- name: CreateChannelRequestHtlc :one
INSERT INTO channel_request_htlcs (
    channel_request_id,
    chan_id, 
    htlc_id, 
    is_settled
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: GetChannelRequestHtlc :one
SELECT * FROM channel_request_htlcs
  WHERE channel_request_id = $1;

-- name: GetChannelRequestHtlcByCircuitKey :one
SELECT * FROM channel_request_htlcs
  WHERE chan_id = $1 AND htlc_id = $2;

-- name: UpdateChannelRequestHtlc :one
UPDATE channel_request_htlcs SET 
    is_settled = $2
  WHERE id = $1
  RETURNING *;
