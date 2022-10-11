// Code generated by sqlc. DO NOT EDIT.
// source: cdr.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createCdr = `-- name: CreateCdr :one
INSERT INTO cdrs (
    uid,
    credential_id,
    country_code, 
    party_id, 
    authorization_id,
    start_date_time,
    stop_date_time,
    auth_id,
    auth_method,
    user_id,
    token_id,
    location_id,
    evse_id,
    connector_id,
    meter_id,
    currency,
    calibration_id,
    total_cost,
    total_energy,
    total_time,
    total_parking_time,
    remark,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
  RETURNING id, uid, credential_id, country_code, party_id, authorization_id, start_date_time, stop_date_time, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, calibration_id, total_cost, total_energy, total_time, total_parking_time, remark, last_updated
`

type CreateCdrParams struct {
	Uid              string          `db:"uid" json:"uid"`
	CredentialID     int64           `db:"credential_id" json:"credentialID"`
	CountryCode      sql.NullString  `db:"country_code" json:"countryCode"`
	PartyID          sql.NullString  `db:"party_id" json:"partyID"`
	AuthorizationID  sql.NullString  `db:"authorization_id" json:"authorizationID"`
	StartDateTime    time.Time       `db:"start_date_time" json:"startDateTime"`
	StopDateTime     sql.NullTime    `db:"stop_date_time" json:"stopDateTime"`
	AuthID           string          `db:"auth_id" json:"authID"`
	AuthMethod       AuthMethodType  `db:"auth_method" json:"authMethod"`
	UserID           int64           `db:"user_id" json:"userID"`
	TokenID          int64           `db:"token_id" json:"tokenID"`
	LocationID       int64           `db:"location_id" json:"locationID"`
	EvseID           int64           `db:"evse_id" json:"evseID"`
	ConnectorID      int64           `db:"connector_id" json:"connectorID"`
	MeterID          sql.NullString  `db:"meter_id" json:"meterID"`
	Currency         string          `db:"currency" json:"currency"`
	CalibrationID    sql.NullInt64   `db:"calibration_id" json:"calibrationID"`
	TotalCost        float64         `db:"total_cost" json:"totalCost"`
	TotalEnergy      float64         `db:"total_energy" json:"totalEnergy"`
	TotalTime        float64         `db:"total_time" json:"totalTime"`
	TotalParkingTime sql.NullFloat64 `db:"total_parking_time" json:"totalParkingTime"`
	Remark           sql.NullString  `db:"remark" json:"remark"`
	LastUpdated      time.Time       `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateCdr(ctx context.Context, arg CreateCdrParams) (Cdr, error) {
	row := q.db.QueryRowContext(ctx, createCdr,
		arg.Uid,
		arg.CredentialID,
		arg.CountryCode,
		arg.PartyID,
		arg.AuthorizationID,
		arg.StartDateTime,
		arg.StopDateTime,
		arg.AuthID,
		arg.AuthMethod,
		arg.UserID,
		arg.TokenID,
		arg.LocationID,
		arg.EvseID,
		arg.ConnectorID,
		arg.MeterID,
		arg.Currency,
		arg.CalibrationID,
		arg.TotalCost,
		arg.TotalEnergy,
		arg.TotalTime,
		arg.TotalParkingTime,
		arg.Remark,
		arg.LastUpdated,
	)
	var i Cdr
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDateTime,
		&i.StopDateTime,
		&i.AuthID,
		&i.AuthMethod,
		&i.UserID,
		&i.TokenID,
		&i.LocationID,
		&i.EvseID,
		&i.ConnectorID,
		&i.MeterID,
		&i.Currency,
		&i.CalibrationID,
		&i.TotalCost,
		&i.TotalEnergy,
		&i.TotalTime,
		&i.TotalParkingTime,
		&i.Remark,
		&i.LastUpdated,
	)
	return i, err
}

const getCdrByLastUpdated = `-- name: GetCdrByLastUpdated :one
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_date_time, stop_date_time, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, calibration_id, total_cost, total_energy, total_time, total_parking_time, remark, last_updated FROM cdrs
  WHERE ($1::BIGINT = -1 OR $1::BIGINT = credential_id) AND
    ($2::TEXT = '' OR $2::TEXT = country_code) AND
    ($3::TEXT = '' OR $3::TEXT = party_id)
  ORDER BY last_updated DESC
  LIMIT 1
`

type GetCdrByLastUpdatedParams struct {
	CredentialID int64  `db:"credential_id" json:"credentialID"`
	CountryCode  string `db:"country_code" json:"countryCode"`
	PartyID      string `db:"party_id" json:"partyID"`
}

func (q *Queries) GetCdrByLastUpdated(ctx context.Context, arg GetCdrByLastUpdatedParams) (Cdr, error) {
	row := q.db.QueryRowContext(ctx, getCdrByLastUpdated, arg.CredentialID, arg.CountryCode, arg.PartyID)
	var i Cdr
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDateTime,
		&i.StopDateTime,
		&i.AuthID,
		&i.AuthMethod,
		&i.UserID,
		&i.TokenID,
		&i.LocationID,
		&i.EvseID,
		&i.ConnectorID,
		&i.MeterID,
		&i.Currency,
		&i.CalibrationID,
		&i.TotalCost,
		&i.TotalEnergy,
		&i.TotalTime,
		&i.TotalParkingTime,
		&i.Remark,
		&i.LastUpdated,
	)
	return i, err
}

const getCdrByUid = `-- name: GetCdrByUid :one
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_date_time, stop_date_time, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, calibration_id, total_cost, total_energy, total_time, total_parking_time, remark, last_updated FROM cdrs
  WHERE uid = $1
`

func (q *Queries) GetCdrByUid(ctx context.Context, uid string) (Cdr, error) {
	row := q.db.QueryRowContext(ctx, getCdrByUid, uid)
	var i Cdr
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDateTime,
		&i.StopDateTime,
		&i.AuthID,
		&i.AuthMethod,
		&i.UserID,
		&i.TokenID,
		&i.LocationID,
		&i.EvseID,
		&i.ConnectorID,
		&i.MeterID,
		&i.Currency,
		&i.CalibrationID,
		&i.TotalCost,
		&i.TotalEnergy,
		&i.TotalTime,
		&i.TotalParkingTime,
		&i.Remark,
		&i.LastUpdated,
	)
	return i, err
}

const listCdrsBySessionStatus = `-- name: ListCdrsBySessionStatus :many
SELECT c.id, c.uid, c.credential_id, c.country_code, c.party_id, c.authorization_id, c.start_date_time, c.stop_date_time, c.auth_id, c.auth_method, c.user_id, c.token_id, c.location_id, c.evse_id, c.connector_id, c.meter_id, c.currency, c.calibration_id, c.total_cost, c.total_energy, c.total_time, c.total_parking_time, c.remark, c.last_updated FROM cdrs c
  INNER JOIN sessions s ON s.authorization_id =cp.authorization_id
  INNER JOIN users u ON u.id = s.user_id
  WHERE u.node_id = $1::BIGINT AND s.status in ($2::session_status_type[])
  ORDER BY c.id
`

type ListCdrsBySessionStatusParams struct {
	NodeID   int64               `db:"node_id" json:"nodeID"`
	Statuses []SessionStatusType `db:"statuses" json:"statuses"`
}

func (q *Queries) ListCdrsBySessionStatus(ctx context.Context, arg ListCdrsBySessionStatusParams) ([]Cdr, error) {
	rows, err := q.db.QueryContext(ctx, listCdrsBySessionStatus, arg.NodeID, pq.Array(arg.Statuses))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cdr
	for rows.Next() {
		var i Cdr
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.CredentialID,
			&i.CountryCode,
			&i.PartyID,
			&i.AuthorizationID,
			&i.StartDateTime,
			&i.StopDateTime,
			&i.AuthID,
			&i.AuthMethod,
			&i.UserID,
			&i.TokenID,
			&i.LocationID,
			&i.EvseID,
			&i.ConnectorID,
			&i.MeterID,
			&i.Currency,
			&i.CalibrationID,
			&i.TotalCost,
			&i.TotalEnergy,
			&i.TotalTime,
			&i.TotalParkingTime,
			&i.Remark,
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
