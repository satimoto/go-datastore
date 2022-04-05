// Code generated by sqlc. DO NOT EDIT.
// source: element_restriction.sql

package db

import (
	"context"
	"database/sql"
)

const createElementRestriction = `-- name: CreateElementRestriction :one
INSERT INTO element_restrictions (
    start_time,
    end_time,
    start_date,
    end_date,
    min_kwh,
    max_kwh,
    min_power,
    max_power,
    min_duration,
    max_duration
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
  RETURNING id, start_time, end_time, start_date, end_date, min_kwh, max_kwh, min_power, max_power, min_duration, max_duration
`

type CreateElementRestrictionParams struct {
	StartTime   sql.NullString  `db:"start_time" json:"startTime"`
	EndTime     sql.NullString  `db:"end_time" json:"endTime"`
	StartDate   sql.NullString  `db:"start_date" json:"startDate"`
	EndDate     sql.NullString  `db:"end_date" json:"endDate"`
	MinKwh      sql.NullFloat64 `db:"min_kwh" json:"minKwh"`
	MaxKwh      sql.NullFloat64 `db:"max_kwh" json:"maxKwh"`
	MinPower    sql.NullFloat64 `db:"min_power" json:"minPower"`
	MaxPower    sql.NullFloat64 `db:"max_power" json:"maxPower"`
	MinDuration sql.NullInt32   `db:"min_duration" json:"minDuration"`
	MaxDuration sql.NullInt32   `db:"max_duration" json:"maxDuration"`
}

func (q *Queries) CreateElementRestriction(ctx context.Context, arg CreateElementRestrictionParams) (ElementRestriction, error) {
	row := q.db.QueryRowContext(ctx, createElementRestriction,
		arg.StartTime,
		arg.EndTime,
		arg.StartDate,
		arg.EndDate,
		arg.MinKwh,
		arg.MaxKwh,
		arg.MinPower,
		arg.MaxPower,
		arg.MinDuration,
		arg.MaxDuration,
	)
	var i ElementRestriction
	err := row.Scan(
		&i.ID,
		&i.StartTime,
		&i.EndTime,
		&i.StartDate,
		&i.EndDate,
		&i.MinKwh,
		&i.MaxKwh,
		&i.MinPower,
		&i.MaxPower,
		&i.MinDuration,
		&i.MaxDuration,
	)
	return i, err
}

const deleteElementRestrictions = `-- name: DeleteElementRestrictions :exec
DELETE FROM element_restrictions r
  USING elements e
  WHERE r.id == e.restriction_id AND e.tariff_id == $1
`

func (q *Queries) DeleteElementRestrictions(ctx context.Context, tariffID int64) error {
	_, err := q.db.ExecContext(ctx, deleteElementRestrictions, tariffID)
	return err
}

const getElementRestriction = `-- name: GetElementRestriction :one
SELECT id, start_time, end_time, start_date, end_date, min_kwh, max_kwh, min_power, max_power, min_duration, max_duration FROM element_restrictions
  WHERE id = $1
`

func (q *Queries) GetElementRestriction(ctx context.Context, id int64) (ElementRestriction, error) {
	row := q.db.QueryRowContext(ctx, getElementRestriction, id)
	var i ElementRestriction
	err := row.Scan(
		&i.ID,
		&i.StartTime,
		&i.EndTime,
		&i.StartDate,
		&i.EndDate,
		&i.MinKwh,
		&i.MaxKwh,
		&i.MinPower,
		&i.MaxPower,
		&i.MinDuration,
		&i.MaxDuration,
	)
	return i, err
}

const updateElementRestriction = `-- name: UpdateElementRestriction :one
UPDATE element_restrictions SET (
    start_time,
    end_time,
    start_date,
    end_date,
    min_kwh,
    max_kwh,
    min_power,
    max_power,
    min_duration,
    max_duration
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE id = $1
  RETURNING id, start_time, end_time, start_date, end_date, min_kwh, max_kwh, min_power, max_power, min_duration, max_duration
`

type UpdateElementRestrictionParams struct {
	ID          int64           `db:"id" json:"id"`
	StartTime   sql.NullString  `db:"start_time" json:"startTime"`
	EndTime     sql.NullString  `db:"end_time" json:"endTime"`
	StartDate   sql.NullString  `db:"start_date" json:"startDate"`
	EndDate     sql.NullString  `db:"end_date" json:"endDate"`
	MinKwh      sql.NullFloat64 `db:"min_kwh" json:"minKwh"`
	MaxKwh      sql.NullFloat64 `db:"max_kwh" json:"maxKwh"`
	MinPower    sql.NullFloat64 `db:"min_power" json:"minPower"`
	MaxPower    sql.NullFloat64 `db:"max_power" json:"maxPower"`
	MinDuration sql.NullInt32   `db:"min_duration" json:"minDuration"`
	MaxDuration sql.NullInt32   `db:"max_duration" json:"maxDuration"`
}

func (q *Queries) UpdateElementRestriction(ctx context.Context, arg UpdateElementRestrictionParams) (ElementRestriction, error) {
	row := q.db.QueryRowContext(ctx, updateElementRestriction,
		arg.ID,
		arg.StartTime,
		arg.EndTime,
		arg.StartDate,
		arg.EndDate,
		arg.MinKwh,
		arg.MaxKwh,
		arg.MinPower,
		arg.MaxPower,
		arg.MinDuration,
		arg.MaxDuration,
	)
	var i ElementRestriction
	err := row.Scan(
		&i.ID,
		&i.StartTime,
		&i.EndTime,
		&i.StartDate,
		&i.EndDate,
		&i.MinKwh,
		&i.MaxKwh,
		&i.MinPower,
		&i.MaxPower,
		&i.MinDuration,
		&i.MaxDuration,
	)
	return i, err
}
