-- name: CreateRoutingEvent :one
INSERT INTO routing_events (
    incoming_chan_id,
    incoming_htlc_id,
    incoming_fiat,
    incoming_msat,
    outgoing_chan_id,
    outgoing_htlc_id,
    outgoing_fiat,
    outgoing_msat,
    fee_fiat,
    fee_msat,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  RETURNING *;
