// Code generated by sqlc. DO NOT EDIT.
// source: session.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
    uid,
    country_code,
    party_id,
    authorization_id,
    start_datetime,
    end_datetime,
    kwh,
    auth_id,
    auth_method,
    location_id,
    meter_id,
    currency,
    total_cost,
    status,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
  RETURNING id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, location_id, meter_id, currency, total_cost, status, last_updated
`

type CreateSessionParams struct {
	Uid             string            `db:"uid" json:"uid"`
	CountryCode     sql.NullString    `db:"country_code" json:"countryCode"`
	PartyID         sql.NullString    `db:"party_id" json:"partyID"`
	AuthorizationID sql.NullString    `db:"authorization_id" json:"authorizationID"`
	StartDatetime   time.Time         `db:"start_datetime" json:"startDatetime"`
	EndDatetime     sql.NullTime      `db:"end_datetime" json:"endDatetime"`
	Kwh             float64           `db:"kwh" json:"kwh"`
	AuthID          string            `db:"auth_id" json:"authID"`
	AuthMethod      AuthMethodType    `db:"auth_method" json:"authMethod"`
	LocationID      int64             `db:"location_id" json:"locationID"`
	MeterID         sql.NullString    `db:"meter_id" json:"meterID"`
	Currency        string            `db:"currency" json:"currency"`
	TotalCost       sql.NullFloat64   `db:"total_cost" json:"totalCost"`
	Status          SessionStatusType `db:"status" json:"status"`
	LastUpdated     time.Time         `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.Uid,
		arg.CountryCode,
		arg.PartyID,
		arg.AuthorizationID,
		arg.StartDatetime,
		arg.EndDatetime,
		arg.Kwh,
		arg.AuthID,
		arg.AuthMethod,
		arg.LocationID,
		arg.MeterID,
		arg.Currency,
		arg.TotalCost,
		arg.Status,
		arg.LastUpdated,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.LocationID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
	)
	return i, err
}

const getSessionByIdentityOrderByLastUpdated = `-- name: GetSessionByIdentityOrderByLastUpdated :one
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, location_id, meter_id, currency, total_cost, status, last_updated FROM sessions
  WHERE country_code = $1 AND party_id = $2
  ORDER BY last_updated DESC
  LIMIT 1
`

type GetSessionByIdentityOrderByLastUpdatedParams struct {
	CountryCode sql.NullString `db:"country_code" json:"countryCode"`
	PartyID     sql.NullString `db:"party_id" json:"partyID"`
}

func (q *Queries) GetSessionByIdentityOrderByLastUpdated(ctx context.Context, arg GetSessionByIdentityOrderByLastUpdatedParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByIdentityOrderByLastUpdated, arg.CountryCode, arg.PartyID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.LocationID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
	)
	return i, err
}

const getSessionByUid = `-- name: GetSessionByUid :one
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, location_id, meter_id, currency, total_cost, status, last_updated FROM sessions
  WHERE uid = $1
`

func (q *Queries) GetSessionByUid(ctx context.Context, uid string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByUid, uid)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.LocationID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
	)
	return i, err
}

const updateSessionByUid = `-- name: UpdateSessionByUid :one
UPDATE sessions SET (
    country_code,
    party_id,
    authorization_id,
    start_datetime,
    end_datetime,
    kwh,
    auth_id,
    auth_method,
    location_id,
    meter_id,
    currency,
    total_cost,
    status,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
  WHERE uid = $1
  RETURNING id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, location_id, meter_id, currency, total_cost, status, last_updated
`

type UpdateSessionByUidParams struct {
	Uid             string            `db:"uid" json:"uid"`
	CountryCode     sql.NullString    `db:"country_code" json:"countryCode"`
	PartyID         sql.NullString    `db:"party_id" json:"partyID"`
	AuthorizationID sql.NullString    `db:"authorization_id" json:"authorizationID"`
	StartDatetime   time.Time         `db:"start_datetime" json:"startDatetime"`
	EndDatetime     sql.NullTime      `db:"end_datetime" json:"endDatetime"`
	Kwh             float64           `db:"kwh" json:"kwh"`
	AuthID          string            `db:"auth_id" json:"authID"`
	AuthMethod      AuthMethodType    `db:"auth_method" json:"authMethod"`
	LocationID      int64             `db:"location_id" json:"locationID"`
	MeterID         sql.NullString    `db:"meter_id" json:"meterID"`
	Currency        string            `db:"currency" json:"currency"`
	TotalCost       sql.NullFloat64   `db:"total_cost" json:"totalCost"`
	Status          SessionStatusType `db:"status" json:"status"`
	LastUpdated     time.Time         `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateSessionByUid(ctx context.Context, arg UpdateSessionByUidParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, updateSessionByUid,
		arg.Uid,
		arg.CountryCode,
		arg.PartyID,
		arg.AuthorizationID,
		arg.StartDatetime,
		arg.EndDatetime,
		arg.Kwh,
		arg.AuthID,
		arg.AuthMethod,
		arg.LocationID,
		arg.MeterID,
		arg.Currency,
		arg.TotalCost,
		arg.Status,
		arg.LastUpdated,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.LocationID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
	)
	return i, err
}
