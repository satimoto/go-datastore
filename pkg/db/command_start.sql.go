// Code generated by sqlc. DO NOT EDIT.
// source: command_start.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createCommandStart = `-- name: CreateCommandStart :one
INSERT INTO command_starts (
    status,
    token_id,
    authorization_id,
    location_id,
    evse_uid,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING id, status, token_id, authorization_id, location_id, evse_uid, last_updated
`

type CreateCommandStartParams struct {
	Status          CommandResponseType `db:"status" json:"status"`
	TokenID         int64               `db:"token_id" json:"tokenID"`
	AuthorizationID sql.NullString      `db:"authorization_id" json:"authorizationID"`
	LocationID      string              `db:"location_id" json:"locationID"`
	EvseUid         sql.NullString      `db:"evse_uid" json:"evseUid"`
	LastUpdated     time.Time           `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateCommandStart(ctx context.Context, arg CreateCommandStartParams) (CommandStart, error) {
	row := q.db.QueryRowContext(ctx, createCommandStart,
		arg.Status,
		arg.TokenID,
		arg.AuthorizationID,
		arg.LocationID,
		arg.EvseUid,
		arg.LastUpdated,
	)
	var i CommandStart
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TokenID,
		&i.AuthorizationID,
		&i.LocationID,
		&i.EvseUid,
		&i.LastUpdated,
	)
	return i, err
}

const getCommandStart = `-- name: GetCommandStart :one
SELECT id, status, token_id, authorization_id, location_id, evse_uid, last_updated FROM command_starts
  WHERE id = $1
`

func (q *Queries) GetCommandStart(ctx context.Context, id int64) (CommandStart, error) {
	row := q.db.QueryRowContext(ctx, getCommandStart, id)
	var i CommandStart
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TokenID,
		&i.AuthorizationID,
		&i.LocationID,
		&i.EvseUid,
		&i.LastUpdated,
	)
	return i, err
}

const getCommandStartByAuthorizationID = `-- name: GetCommandStartByAuthorizationID :one
SELECT id, status, token_id, authorization_id, location_id, evse_uid, last_updated FROM command_starts
  WHERE authorization_id = $1
`

func (q *Queries) GetCommandStartByAuthorizationID(ctx context.Context, authorizationID sql.NullString) (CommandStart, error) {
	row := q.db.QueryRowContext(ctx, getCommandStartByAuthorizationID, authorizationID)
	var i CommandStart
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TokenID,
		&i.AuthorizationID,
		&i.LocationID,
		&i.EvseUid,
		&i.LastUpdated,
	)
	return i, err
}

const updateCommandStart = `-- name: UpdateCommandStart :one
UPDATE command_starts SET (
    status,
    last_updated
  ) = ($2, $3)
  WHERE id = $1
  RETURNING id, status, token_id, authorization_id, location_id, evse_uid, last_updated
`

type UpdateCommandStartParams struct {
	ID          int64               `db:"id" json:"id"`
	Status      CommandResponseType `db:"status" json:"status"`
	LastUpdated time.Time           `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateCommandStart(ctx context.Context, arg UpdateCommandStartParams) (CommandStart, error) {
	row := q.db.QueryRowContext(ctx, updateCommandStart, arg.ID, arg.Status, arg.LastUpdated)
	var i CommandStart
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TokenID,
		&i.AuthorizationID,
		&i.LocationID,
		&i.EvseUid,
		&i.LastUpdated,
	)
	return i, err
}
