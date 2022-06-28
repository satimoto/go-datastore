// Code generated by sqlc. DO NOT EDIT.
// source: routing_event.sql

package db

import (
	"context"
	"time"
)

const createRoutingEvent = `-- name: CreateRoutingEvent :one
INSERT INTO routing_events (
    node_id,
    event_type,
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
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
  RETURNING id, node_id, event_type, currency, currency_rate, currency_rate_msat, incoming_chan_id, incoming_htlc_id, incoming_fiat, incoming_msat, outgoing_chan_id, outgoing_htlc_id, outgoing_fiat, outgoing_msat, fee_fiat, fee_msat, last_updated
`

type CreateRoutingEventParams struct {
	NodeID           int64            `db:"node_id" json:"nodeID"`
	EventType        RoutingEventType `db:"event_type" json:"eventType"`
	Currency         string           `db:"currency" json:"currency"`
	CurrencyRate     int64            `db:"currency_rate" json:"currencyRate"`
	CurrencyRateMsat int64            `db:"currency_rate_msat" json:"currencyRateMsat"`
	IncomingChanID   int64            `db:"incoming_chan_id" json:"incomingChanID"`
	IncomingHtlcID   int64            `db:"incoming_htlc_id" json:"incomingHtlcID"`
	IncomingFiat     float64          `db:"incoming_fiat" json:"incomingFiat"`
	IncomingMsat     int64            `db:"incoming_msat" json:"incomingMsat"`
	OutgoingChanID   int64            `db:"outgoing_chan_id" json:"outgoingChanID"`
	OutgoingHtlcID   int64            `db:"outgoing_htlc_id" json:"outgoingHtlcID"`
	OutgoingFiat     float64          `db:"outgoing_fiat" json:"outgoingFiat"`
	OutgoingMsat     int64            `db:"outgoing_msat" json:"outgoingMsat"`
	FeeFiat          float64          `db:"fee_fiat" json:"feeFiat"`
	FeeMsat          int64            `db:"fee_msat" json:"feeMsat"`
	LastUpdated      time.Time        `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateRoutingEvent(ctx context.Context, arg CreateRoutingEventParams) (RoutingEvent, error) {
	row := q.db.QueryRowContext(ctx, createRoutingEvent,
		arg.NodeID,
		arg.EventType,
		arg.Currency,
		arg.CurrencyRate,
		arg.CurrencyRateMsat,
		arg.IncomingChanID,
		arg.IncomingHtlcID,
		arg.IncomingFiat,
		arg.IncomingMsat,
		arg.OutgoingChanID,
		arg.OutgoingHtlcID,
		arg.OutgoingFiat,
		arg.OutgoingMsat,
		arg.FeeFiat,
		arg.FeeMsat,
		arg.LastUpdated,
	)
	var i RoutingEvent
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.EventType,
		&i.Currency,
		&i.CurrencyRate,
		&i.CurrencyRateMsat,
		&i.IncomingChanID,
		&i.IncomingHtlcID,
		&i.IncomingFiat,
		&i.IncomingMsat,
		&i.OutgoingChanID,
		&i.OutgoingHtlcID,
		&i.OutgoingFiat,
		&i.OutgoingMsat,
		&i.FeeFiat,
		&i.FeeMsat,
		&i.LastUpdated,
	)
	return i, err
}
