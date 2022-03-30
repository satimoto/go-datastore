// Code generated by sqlc. DO NOT EDIT.
// source: command_stop.sql

package db

import (
	"context"
)

const createCommandStop = `-- name: CreateCommandStop :one
INSERT INTO command_stops (
    status,
    session_id
  ) VALUES ($1, $2)
  RETURNING id, status, session_id
`

type CreateCommandStopParams struct {
	Status    CommandResponseType `db:"status" json:"status"`
	SessionID string              `db:"session_id" json:"sessionID"`
}

func (q *Queries) CreateCommandStop(ctx context.Context, arg CreateCommandStopParams) (CommandStop, error) {
	row := q.db.QueryRowContext(ctx, createCommandStop, arg.Status, arg.SessionID)
	var i CommandStop
	err := row.Scan(&i.ID, &i.Status, &i.SessionID)
	return i, err
}

const updateCommandStop = `-- name: UpdateCommandStop :one
UPDATE command_stops SET status = $2
  WHERE id = $1
  RETURNING id, status, session_id
`

type UpdateCommandStopParams struct {
	ID     int64               `db:"id" json:"id"`
	Status CommandResponseType `db:"status" json:"status"`
}

func (q *Queries) UpdateCommandStop(ctx context.Context, arg UpdateCommandStopParams) (CommandStop, error) {
	row := q.db.QueryRowContext(ctx, updateCommandStop, arg.ID, arg.Status)
	var i CommandStop
	err := row.Scan(&i.ID, &i.Status, &i.SessionID)
	return i, err
}
