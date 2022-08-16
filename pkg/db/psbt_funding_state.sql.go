// Code generated by sqlc. DO NOT EDIT.
// source: psbt_funding_state.sql

package db

import (
	"context"
	"time"
)

const createPsbtFundingState = `-- name: CreatePsbtFundingState :one
INSERT INTO psbt_funding_states (
    node_id,
    base_psbt,
    psbt, 
    expiry_date
  ) VALUES ($1, $2, $3, $4)
  RETURNING id, node_id, base_psbt, psbt, funded_psbt, expiry_date
`

type CreatePsbtFundingStateParams struct {
	NodeID     int64     `db:"node_id" json:"nodeID"`
	BasePsbt   []byte    `db:"base_psbt" json:"basePsbt"`
	Psbt       []byte    `db:"psbt" json:"psbt"`
	ExpiryDate time.Time `db:"expiry_date" json:"expiryDate"`
}

func (q *Queries) CreatePsbtFundingState(ctx context.Context, arg CreatePsbtFundingStateParams) (PsbtFundingState, error) {
	row := q.db.QueryRowContext(ctx, createPsbtFundingState,
		arg.NodeID,
		arg.BasePsbt,
		arg.Psbt,
		arg.ExpiryDate,
	)
	var i PsbtFundingState
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.BasePsbt,
		&i.Psbt,
		&i.FundedPsbt,
		&i.ExpiryDate,
	)
	return i, err
}

const getPsbtFundingState = `-- name: GetPsbtFundingState :one
SELECT id, node_id, base_psbt, psbt, funded_psbt, expiry_date FROM psbt_funding_states
  WHERE node_id = $1 AND funded_psbt is null
`

func (q *Queries) GetPsbtFundingState(ctx context.Context, nodeID int64) (PsbtFundingState, error) {
	row := q.db.QueryRowContext(ctx, getPsbtFundingState, nodeID)
	var i PsbtFundingState
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.BasePsbt,
		&i.Psbt,
		&i.FundedPsbt,
		&i.ExpiryDate,
	)
	return i, err
}

const listPsbtFundingStates = `-- name: ListPsbtFundingStates :many
SELECT id, node_id, base_psbt, psbt, funded_psbt, expiry_date FROM psbt_funding_states
  WHERE node_id = $1 AND funded_psbt is null
  ORDER BY id
`

func (q *Queries) ListPsbtFundingStates(ctx context.Context, nodeID int64) ([]PsbtFundingState, error) {
	rows, err := q.db.QueryContext(ctx, listPsbtFundingStates, nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PsbtFundingState
	for rows.Next() {
		var i PsbtFundingState
		if err := rows.Scan(
			&i.ID,
			&i.NodeID,
			&i.BasePsbt,
			&i.Psbt,
			&i.FundedPsbt,
			&i.ExpiryDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePsbtFundingState = `-- name: UpdatePsbtFundingState :one
UPDATE psbt_funding_states SET (
    psbt,
    funded_psbt
  ) = ($2, $3)
  WHERE id = $1
  RETURNING id, node_id, base_psbt, psbt, funded_psbt, expiry_date
`

type UpdatePsbtFundingStateParams struct {
	ID         int64  `db:"id" json:"id"`
	Psbt       []byte `db:"psbt" json:"psbt"`
	FundedPsbt []byte `db:"funded_psbt" json:"fundedPsbt"`
}

func (q *Queries) UpdatePsbtFundingState(ctx context.Context, arg UpdatePsbtFundingStateParams) (PsbtFundingState, error) {
	row := q.db.QueryRowContext(ctx, updatePsbtFundingState, arg.ID, arg.Psbt, arg.FundedPsbt)
	var i PsbtFundingState
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.BasePsbt,
		&i.Psbt,
		&i.FundedPsbt,
		&i.ExpiryDate,
	)
	return i, err
}
