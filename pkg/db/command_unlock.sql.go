// Code generated by sqlc. DO NOT EDIT.
// source: command_unlock.sql

package db

import (
	"context"
	"time"
)

const createCommandUnlock = `-- name: CreateCommandUnlock :one
INSERT INTO command_unlocks (
    status,
    location_id,
    evse_uid,
    connector_id,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING id, status, location_id, evse_uid, connector_id, last_updated
`

type CreateCommandUnlockParams struct {
	Status      CommandResponseType `db:"status" json:"status"`
	LocationID  string              `db:"location_id" json:"locationID"`
	EvseUid     string              `db:"evse_uid" json:"evseUid"`
	ConnectorID string              `db:"connector_id" json:"connectorID"`
	LastUpdated time.Time           `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateCommandUnlock(ctx context.Context, arg CreateCommandUnlockParams) (CommandUnlock, error) {
	row := q.db.QueryRowContext(ctx, createCommandUnlock,
		arg.Status,
		arg.LocationID,
		arg.EvseUid,
		arg.ConnectorID,
		arg.LastUpdated,
	)
	var i CommandUnlock
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.LocationID,
		&i.EvseUid,
		&i.ConnectorID,
		&i.LastUpdated,
	)
	return i, err
}

const getCommandUnlock = `-- name: GetCommandUnlock :one
SELECT id, status, location_id, evse_uid, connector_id, last_updated FROM command_unlocks
  WHERE id = $1
`

func (q *Queries) GetCommandUnlock(ctx context.Context, id int64) (CommandUnlock, error) {
	row := q.db.QueryRowContext(ctx, getCommandUnlock, id)
	var i CommandUnlock
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.LocationID,
		&i.EvseUid,
		&i.ConnectorID,
		&i.LastUpdated,
	)
	return i, err
}

const updateCommandUnlock = `-- name: UpdateCommandUnlock :one
UPDATE command_unlocks SET (
    status,
    last_updated
  ) = ($2, $3)
  WHERE id = $1
  RETURNING id, status, location_id, evse_uid, connector_id, last_updated
`

type UpdateCommandUnlockParams struct {
	ID          int64               `db:"id" json:"id"`
	Status      CommandResponseType `db:"status" json:"status"`
	LastUpdated time.Time           `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateCommandUnlock(ctx context.Context, arg UpdateCommandUnlockParams) (CommandUnlock, error) {
	row := q.db.QueryRowContext(ctx, updateCommandUnlock, arg.ID, arg.Status, arg.LastUpdated)
	var i CommandUnlock
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.LocationID,
		&i.EvseUid,
		&i.ConnectorID,
		&i.LastUpdated,
	)
	return i, err
}
