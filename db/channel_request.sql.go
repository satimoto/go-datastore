// Code generated by sqlc. DO NOT EDIT.
// source: channel_request.sql

package db

import (
	"context"
	"database/sql"
)

const createChannelRequest = `-- name: CreateChannelRequest :one
INSERT INTO channel_requests (
    status,
    pubkey, 
    preimage, 
    payment_hash, 
    payment_addr,
    amount_msat,
    settled_msat
  ) VALUES ($1, $2, $3, $4, $5, $6, $7)
  RETURNING id, status, pubkey, preimage, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id, output_index
`

type CreateChannelRequestParams struct {
	Status      ChannelRequestStatus `db:"status" json:"status"`
	Pubkey      []byte               `db:"pubkey" json:"pubkey"`
	Preimage    []byte               `db:"preimage" json:"preimage"`
	PaymentHash []byte               `db:"payment_hash" json:"paymentHash"`
	PaymentAddr []byte               `db:"payment_addr" json:"paymentAddr"`
	AmountMsat  int64                `db:"amount_msat" json:"amountMsat"`
	SettledMsat int64                `db:"settled_msat" json:"settledMsat"`
}

func (q *Queries) CreateChannelRequest(ctx context.Context, arg CreateChannelRequestParams) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, createChannelRequest,
		arg.Status,
		arg.Pubkey,
		arg.Preimage,
		arg.PaymentHash,
		arg.PaymentAddr,
		arg.AmountMsat,
		arg.SettledMsat,
	)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.Pubkey,
		&i.Preimage,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxID,
		&i.OutputIndex,
	)
	return i, err
}

const deleteChannelRequest = `-- name: DeleteChannelRequest :exec
DELETE FROM channel_requests
  WHERE id = $1
`

func (q *Queries) DeleteChannelRequest(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteChannelRequest, id)
	return err
}

const getChannelRequest = `-- name: GetChannelRequest :one
SELECT id, status, pubkey, preimage, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id, output_index FROM channel_requests
  WHERE id = $1
`

func (q *Queries) GetChannelRequest(ctx context.Context, id int64) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, getChannelRequest, id)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.Pubkey,
		&i.Preimage,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxID,
		&i.OutputIndex,
	)
	return i, err
}

const getChannelRequestByPaymentHash = `-- name: GetChannelRequestByPaymentHash :one
SELECT id, status, pubkey, preimage, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id, output_index FROM channel_requests
  WHERE payment_hash = $1 OR sha256('probing-01:' || payment_hash) = $1
`

func (q *Queries) GetChannelRequestByPaymentHash(ctx context.Context, paymentHash []byte) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, getChannelRequestByPaymentHash, paymentHash)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.Pubkey,
		&i.Preimage,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxID,
		&i.OutputIndex,
	)
	return i, err
}

const updateChannelRequest = `-- name: UpdateChannelRequest :one
UPDATE channel_requests SET (
    status,
    settled_msat,
    funding_tx_id, 
    output_index
  ) = ($2, $3, $4, $5)
  WHERE id = $1
  RETURNING id, status, pubkey, preimage, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id, output_index
`

type UpdateChannelRequestParams struct {
	ID          int64                `db:"id" json:"id"`
	Status      ChannelRequestStatus `db:"status" json:"status"`
	SettledMsat int64                `db:"settled_msat" json:"settledMsat"`
	FundingTxID []byte               `db:"funding_tx_id" json:"fundingTxID"`
	OutputIndex sql.NullInt64        `db:"output_index" json:"outputIndex"`
}

func (q *Queries) UpdateChannelRequest(ctx context.Context, arg UpdateChannelRequestParams) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, updateChannelRequest,
		arg.ID,
		arg.Status,
		arg.SettledMsat,
		arg.FundingTxID,
		arg.OutputIndex,
	)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.Pubkey,
		&i.Preimage,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxID,
		&i.OutputIndex,
	)
	return i, err
}

const updateChannelRequestByChannelPoint = `-- name: UpdateChannelRequestByChannelPoint :one
UPDATE channel_requests SET status = $2
  WHERE id = $1
  RETURNING id, status, pubkey, preimage, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id, output_index
`

type UpdateChannelRequestByChannelPointParams struct {
	ID     int64                `db:"id" json:"id"`
	Status ChannelRequestStatus `db:"status" json:"status"`
}

func (q *Queries) UpdateChannelRequestByChannelPoint(ctx context.Context, arg UpdateChannelRequestByChannelPointParams) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, updateChannelRequestByChannelPoint, arg.ID, arg.Status)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.Pubkey,
		&i.Preimage,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxID,
		&i.OutputIndex,
	)
	return i, err
}
