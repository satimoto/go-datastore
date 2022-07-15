-- name: CreateRoutingEvent :one
INSERT INTO routing_events (
    node_id,
    event_type,
    event_status,
    currency,
    currency_rate,
    currency_rate_msat,
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
    wire_failure,
    failure_detail,
    failure_string,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
  RETURNING *;

-- name: UpdateRoutingEvent :one
UPDATE routing_events SET (
    event_status,
    wire_failure,
    failure_detail,
    failure_string,
    last_updated
  ) = ($5, $6, $7, $8, $9)
  WHERE incoming_chan_id = $1 AND incoming_htlc_id = $2 AND outgoing_chan_id = $3 AND outgoing_htlc_id = $4
  RETURNING *;
