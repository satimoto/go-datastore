// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: authentication.sql

package db

import (
	"context"
	"database/sql"
)

const createAuthentication = `-- name: CreateAuthentication :one
INSERT INTO authentications (
    code, 
    action, 
    challenge
  ) VALUES ($1, $2, $3)
  RETURNING id, code, action, challenge, signature, linking_pubkey
`

type CreateAuthenticationParams struct {
	Code      string                `db:"code" json:"code"`
	Action    AuthenticationActions `db:"action" json:"action"`
	Challenge string                `db:"challenge" json:"challenge"`
}

func (q *Queries) CreateAuthentication(ctx context.Context, arg CreateAuthenticationParams) (Authentication, error) {
	row := q.db.QueryRowContext(ctx, createAuthentication, arg.Code, arg.Action, arg.Challenge)
	var i Authentication
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Action,
		&i.Challenge,
		&i.Signature,
		&i.LinkingPubkey,
	)
	return i, err
}

const deleteAuthentication = `-- name: DeleteAuthentication :exec
DELETE FROM authentications
  WHERE id = $1
`

func (q *Queries) DeleteAuthentication(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthentication, id)
	return err
}

const getAuthenticationByChallenge = `-- name: GetAuthenticationByChallenge :one
SELECT id, code, action, challenge, signature, linking_pubkey FROM authentications
  WHERE challenge = $1
`

func (q *Queries) GetAuthenticationByChallenge(ctx context.Context, challenge string) (Authentication, error) {
	row := q.db.QueryRowContext(ctx, getAuthenticationByChallenge, challenge)
	var i Authentication
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Action,
		&i.Challenge,
		&i.Signature,
		&i.LinkingPubkey,
	)
	return i, err
}

const getAuthenticationByCode = `-- name: GetAuthenticationByCode :one
SELECT id, code, action, challenge, signature, linking_pubkey FROM authentications
  WHERE code = $1
`

func (q *Queries) GetAuthenticationByCode(ctx context.Context, code string) (Authentication, error) {
	row := q.db.QueryRowContext(ctx, getAuthenticationByCode, code)
	var i Authentication
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Action,
		&i.Challenge,
		&i.Signature,
		&i.LinkingPubkey,
	)
	return i, err
}

const updateAuthentication = `-- name: UpdateAuthentication :one
UPDATE authentications SET (
    signature, 
    linking_pubkey
  ) = ($2, $3)
  WHERE id = $1
  RETURNING id, code, action, challenge, signature, linking_pubkey
`

type UpdateAuthenticationParams struct {
	ID            int64          `db:"id" json:"id"`
	Signature     sql.NullString `db:"signature" json:"signature"`
	LinkingPubkey sql.NullString `db:"linking_pubkey" json:"linkingPubkey"`
}

func (q *Queries) UpdateAuthentication(ctx context.Context, arg UpdateAuthenticationParams) (Authentication, error) {
	row := q.db.QueryRowContext(ctx, updateAuthentication, arg.ID, arg.Signature, arg.LinkingPubkey)
	var i Authentication
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Action,
		&i.Challenge,
		&i.Signature,
		&i.LinkingPubkey,
	)
	return i, err
}
