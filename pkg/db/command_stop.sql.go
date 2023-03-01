// Code generated by sqlc. DO NOT EDIT.
// source: command_stop.sql

package db

import (
	"context"
	"time"
)

const createCommandStop = `-- name: CreateCommandStop :one
INSERT INTO command_stops (
    status,
    session_id,
    last_updated
  ) VALUES ($1, $2, $3)
  RETURNING id, status, session_id, last_updated
`

type CreateCommandStopParams struct {
	Status      CommandResponseType `db:"status" json:"status"`
	SessionID   string              `db:"session_id" json:"sessionID"`
	LastUpdated time.Time           `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateCommandStop(ctx context.Context, arg CreateCommandStopParams) (CommandStop, error) {
	row := q.db.QueryRowContext(ctx, createCommandStop, arg.Status, arg.SessionID, arg.LastUpdated)
	var i CommandStop
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.SessionID,
		&i.LastUpdated,
	)
	return i, err
}

const getCommandStop = `-- name: GetCommandStop :one
SELECT id, status, session_id, last_updated FROM command_stops
  WHERE id = $1
`

func (q *Queries) GetCommandStop(ctx context.Context, id int64) (CommandStop, error) {
	row := q.db.QueryRowContext(ctx, getCommandStop, id)
	var i CommandStop
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.SessionID,
		&i.LastUpdated,
	)
	return i, err
}

const updateCommandStop = `-- name: UpdateCommandStop :one
UPDATE command_stops SET (
    status,
    last_updated
  ) = ($2, $3)
  WHERE id = $1
  RETURNING id, status, session_id, last_updated
`

type UpdateCommandStopParams struct {
	ID          int64               `db:"id" json:"id"`
	Status      CommandResponseType `db:"status" json:"status"`
	LastUpdated time.Time           `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateCommandStop(ctx context.Context, arg UpdateCommandStopParams) (CommandStop, error) {
	row := q.db.QueryRowContext(ctx, updateCommandStop, arg.ID, arg.Status, arg.LastUpdated)
	var i CommandStop
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.SessionID,
		&i.LastUpdated,
	)
	return i, err
}

const updateCommandStopBySessionID = `-- name: UpdateCommandStopBySessionID :one
UPDATE command_stops SET (
    status,
    last_updated
  ) = ($2, $3)
  WHERE session_id = $1 AND status = 'REQUESTED'
  RETURNING id, status, session_id, last_updated
`

type UpdateCommandStopBySessionIDParams struct {
	SessionID   string              `db:"session_id" json:"sessionID"`
	Status      CommandResponseType `db:"status" json:"status"`
	LastUpdated time.Time           `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateCommandStopBySessionID(ctx context.Context, arg UpdateCommandStopBySessionIDParams) (CommandStop, error) {
	row := q.db.QueryRowContext(ctx, updateCommandStopBySessionID, arg.SessionID, arg.Status, arg.LastUpdated)
	var i CommandStop
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.SessionID,
		&i.LastUpdated,
	)
	return i, err
}
