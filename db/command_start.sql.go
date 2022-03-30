// Code generated by sqlc. DO NOT EDIT.
// source: command_start.sql

package db

import (
	"context"
	"database/sql"
)

const createCommandStart = `-- name: CreateCommandStart :one
INSERT INTO command_starts (
    status,
    token_id,
    location_id,
    evse_uid
  ) VALUES ($1, $2, $3, $4)
  RETURNING id, status, token_id, location_id, evse_uid
`

type CreateCommandStartParams struct {
	Status     CommandResponseType `db:"status" json:"status"`
	TokenID    int64               `db:"token_id" json:"tokenID"`
	LocationID string              `db:"location_id" json:"locationID"`
	EvseUid    sql.NullString      `db:"evse_uid" json:"evseUid"`
}

func (q *Queries) CreateCommandStart(ctx context.Context, arg CreateCommandStartParams) (CommandStart, error) {
	row := q.db.QueryRowContext(ctx, createCommandStart,
		arg.Status,
		arg.TokenID,
		arg.LocationID,
		arg.EvseUid,
	)
	var i CommandStart
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TokenID,
		&i.LocationID,
		&i.EvseUid,
	)
	return i, err
}

const updateCommandStart = `-- name: UpdateCommandStart :one
UPDATE command_starts SET status = $2
  WHERE id = $1
  RETURNING id, status, token_id, location_id, evse_uid
`

type UpdateCommandStartParams struct {
	ID     int64               `db:"id" json:"id"`
	Status CommandResponseType `db:"status" json:"status"`
}

func (q *Queries) UpdateCommandStart(ctx context.Context, arg UpdateCommandStartParams) (CommandStart, error) {
	row := q.db.QueryRowContext(ctx, updateCommandStart, arg.ID, arg.Status)
	var i CommandStart
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TokenID,
		&i.LocationID,
		&i.EvseUid,
	)
	return i, err
}
