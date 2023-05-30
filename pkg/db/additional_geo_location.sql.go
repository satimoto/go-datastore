// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: additional_geo_location.sql

package db

import (
	"context"
	"database/sql"
)

const createAdditionalGeoLocation = `-- name: CreateAdditionalGeoLocation :one
INSERT INTO additional_geo_locations (
    location_id,
    display_text_id,
    latitude, 
    latitude_float, 
    longitude, 
    longitude_float 
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING location_id, display_text_id, latitude, latitude_float, longitude, longitude_float
`

type CreateAdditionalGeoLocationParams struct {
	LocationID     int64         `db:"location_id" json:"locationID"`
	DisplayTextID  sql.NullInt64 `db:"display_text_id" json:"displayTextID"`
	Latitude       string        `db:"latitude" json:"latitude"`
	LatitudeFloat  float64       `db:"latitude_float" json:"latitudeFloat"`
	Longitude      string        `db:"longitude" json:"longitude"`
	LongitudeFloat float64       `db:"longitude_float" json:"longitudeFloat"`
}

func (q *Queries) CreateAdditionalGeoLocation(ctx context.Context, arg CreateAdditionalGeoLocationParams) (AdditionalGeoLocation, error) {
	row := q.db.QueryRowContext(ctx, createAdditionalGeoLocation,
		arg.LocationID,
		arg.DisplayTextID,
		arg.Latitude,
		arg.LatitudeFloat,
		arg.Longitude,
		arg.LongitudeFloat,
	)
	var i AdditionalGeoLocation
	err := row.Scan(
		&i.LocationID,
		&i.DisplayTextID,
		&i.Latitude,
		&i.LatitudeFloat,
		&i.Longitude,
		&i.LongitudeFloat,
	)
	return i, err
}

const deleteAdditionalGeoLocations = `-- name: DeleteAdditionalGeoLocations :exec
DELETE FROM additional_geo_locations
  WHERE location_id = $1
`

func (q *Queries) DeleteAdditionalGeoLocations(ctx context.Context, locationID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAdditionalGeoLocations, locationID)
	return err
}

const listAdditionalGeoLocations = `-- name: ListAdditionalGeoLocations :many
SELECT location_id, display_text_id, latitude, latitude_float, longitude, longitude_float FROM additional_geo_locations
  WHERE location_id = $1
`

func (q *Queries) ListAdditionalGeoLocations(ctx context.Context, locationID int64) ([]AdditionalGeoLocation, error) {
	rows, err := q.db.QueryContext(ctx, listAdditionalGeoLocations, locationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AdditionalGeoLocation
	for rows.Next() {
		var i AdditionalGeoLocation
		if err := rows.Scan(
			&i.LocationID,
			&i.DisplayTextID,
			&i.Latitude,
			&i.LatitudeFloat,
			&i.Longitude,
			&i.LongitudeFloat,
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
