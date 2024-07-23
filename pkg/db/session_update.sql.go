// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: session_update.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createSessionUpdate = `-- name: CreateSessionUpdate :one
INSERT INTO session_updates (
    session_id,
    user_id,
    kwh,
    total_cost,
    status,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING id, session_id, user_id, kwh, total_cost, status, last_updated
`

type CreateSessionUpdateParams struct {
	SessionID   int64             `db:"session_id" json:"sessionID"`
	UserID      int64             `db:"user_id" json:"userID"`
	Kwh         float64           `db:"kwh" json:"kwh"`
	TotalCost   sql.NullFloat64   `db:"total_cost" json:"totalCost"`
	Status      SessionStatusType `db:"status" json:"status"`
	LastUpdated time.Time         `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateSessionUpdate(ctx context.Context, arg CreateSessionUpdateParams) (SessionUpdate, error) {
	row := q.db.QueryRowContext(ctx, createSessionUpdate,
		arg.SessionID,
		arg.UserID,
		arg.Kwh,
		arg.TotalCost,
		arg.Status,
		arg.LastUpdated,
	)
	var i SessionUpdate
	err := row.Scan(
		&i.ID,
		&i.SessionID,
		&i.UserID,
		&i.Kwh,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
	)
	return i, err
}

const listSessionUpdatesBySessionID = `-- name: ListSessionUpdatesBySessionID :many
SELECT id, session_id, user_id, kwh, total_cost, status, last_updated FROM session_updates
  WHERE session_id = $1
  ORDER BY id
`

func (q *Queries) ListSessionUpdatesBySessionID(ctx context.Context, sessionID int64) ([]SessionUpdate, error) {
	rows, err := q.db.QueryContext(ctx, listSessionUpdatesBySessionID, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SessionUpdate
	for rows.Next() {
		var i SessionUpdate
		if err := rows.Scan(
			&i.ID,
			&i.SessionID,
			&i.UserID,
			&i.Kwh,
			&i.TotalCost,
			&i.Status,
			&i.LastUpdated,
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
