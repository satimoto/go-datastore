// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: calibration.sql

package db

import (
	"context"
	"database/sql"
)

const createCalibration = `-- name: CreateCalibration :one
INSERT INTO calibrations (
    encoding_method,
    encoding_method_version,
    public_key,
    url
  ) VALUES ($1, $2, $3, $4)
  RETURNING id, encoding_method, encoding_method_version, public_key, url
`

type CreateCalibrationParams struct {
	EncodingMethod        string         `db:"encoding_method" json:"encodingMethod"`
	EncodingMethodVersion sql.NullInt32  `db:"encoding_method_version" json:"encodingMethodVersion"`
	PublicKey             sql.NullString `db:"public_key" json:"publicKey"`
	Url                   sql.NullString `db:"url" json:"url"`
}

func (q *Queries) CreateCalibration(ctx context.Context, arg CreateCalibrationParams) (Calibration, error) {
	row := q.db.QueryRowContext(ctx, createCalibration,
		arg.EncodingMethod,
		arg.EncodingMethodVersion,
		arg.PublicKey,
		arg.Url,
	)
	var i Calibration
	err := row.Scan(
		&i.ID,
		&i.EncodingMethod,
		&i.EncodingMethodVersion,
		&i.PublicKey,
		&i.Url,
	)
	return i, err
}

const getCalibration = `-- name: GetCalibration :one
SELECT id, encoding_method, encoding_method_version, public_key, url FROM calibrations
  WHERE id = $1
`

func (q *Queries) GetCalibration(ctx context.Context, id int64) (Calibration, error) {
	row := q.db.QueryRowContext(ctx, getCalibration, id)
	var i Calibration
	err := row.Scan(
		&i.ID,
		&i.EncodingMethod,
		&i.EncodingMethodVersion,
		&i.PublicKey,
		&i.Url,
	)
	return i, err
}
