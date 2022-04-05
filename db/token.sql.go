// Code generated by sqlc. DO NOT EDIT.
// source: token.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createToken = `-- name: CreateToken :one
INSERT INTO tokens (
    uid, 
    type,
    auth_id,
    visual_number,
    issuer,
    allowed,
    valid,
    whitelist,
    language,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
  RETURNING id, uid, type, auth_id, visual_number, issuer, allowed, valid, whitelist, language, last_updated
`

type CreateTokenParams struct {
	Uid          string             `db:"uid" json:"uid"`
	Type         TokenType          `db:"type" json:"type"`
	AuthID       string             `db:"auth_id" json:"authID"`
	VisualNumber sql.NullString     `db:"visual_number" json:"visualNumber"`
	Issuer       string             `db:"issuer" json:"issuer"`
	Allowed      TokenAllowedType   `db:"allowed" json:"allowed"`
	Valid        bool               `db:"valid" json:"valid"`
	Whitelist    TokenWhitelistType `db:"whitelist" json:"whitelist"`
	Language     sql.NullString     `db:"language" json:"language"`
	LastUpdated  time.Time          `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, createToken,
		arg.Uid,
		arg.Type,
		arg.AuthID,
		arg.VisualNumber,
		arg.Issuer,
		arg.Allowed,
		arg.Valid,
		arg.Whitelist,
		arg.Language,
		arg.LastUpdated,
	)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Type,
		&i.AuthID,
		&i.VisualNumber,
		&i.Issuer,
		&i.Allowed,
		&i.Valid,
		&i.Whitelist,
		&i.Language,
		&i.LastUpdated,
	)
	return i, err
}

const deleteTokenByUid = `-- name: DeleteTokenByUid :exec
DELETE FROM tokens
  WHERE uid = $1
`

func (q *Queries) DeleteTokenByUid(ctx context.Context, uid string) error {
	_, err := q.db.ExecContext(ctx, deleteTokenByUid, uid)
	return err
}

const getToken = `-- name: GetToken :one
SELECT id, uid, type, auth_id, visual_number, issuer, allowed, valid, whitelist, language, last_updated FROM tokens
  WHERE id = $1
`

func (q *Queries) GetToken(ctx context.Context, id int64) (Token, error) {
	row := q.db.QueryRowContext(ctx, getToken, id)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Type,
		&i.AuthID,
		&i.VisualNumber,
		&i.Issuer,
		&i.Allowed,
		&i.Valid,
		&i.Whitelist,
		&i.Language,
		&i.LastUpdated,
	)
	return i, err
}

const getTokenByUid = `-- name: GetTokenByUid :one
SELECT id, uid, type, auth_id, visual_number, issuer, allowed, valid, whitelist, language, last_updated FROM tokens
  WHERE uid = $1
`

func (q *Queries) GetTokenByUid(ctx context.Context, uid string) (Token, error) {
	row := q.db.QueryRowContext(ctx, getTokenByUid, uid)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Type,
		&i.AuthID,
		&i.VisualNumber,
		&i.Issuer,
		&i.Allowed,
		&i.Valid,
		&i.Whitelist,
		&i.Language,
		&i.LastUpdated,
	)
	return i, err
}

const listTokens = `-- name: ListTokens :many
SELECT id, uid, type, auth_id, visual_number, issuer, allowed, valid, whitelist, language, last_updated FROM tokens
  WHERE 
    ($1::text = '' or last_updated > $1::text) and 
    ($2::text = '' or last_updated < $2::text)
  ORDER BY id
  LIMIT $4::bigint
  OFFSET $3::bigint
`

type ListTokensParams struct {
	FilterDateFrom string `db:"filter_date_from" json:"filterDateFrom"`
	FilterDateTo   string `db:"filter_date_to" json:"filterDateTo"`
	FilterOffset   int64  `db:"filter_offset" json:"filterOffset"`
	FilterLimit    int64  `db:"filter_limit" json:"filterLimit"`
}

func (q *Queries) ListTokens(ctx context.Context, arg ListTokensParams) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, listTokens,
		arg.FilterDateFrom,
		arg.FilterDateTo,
		arg.FilterOffset,
		arg.FilterLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.Type,
			&i.AuthID,
			&i.VisualNumber,
			&i.Issuer,
			&i.Allowed,
			&i.Valid,
			&i.Whitelist,
			&i.Language,
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

const updateTokenByUid = `-- name: UpdateTokenByUid :one
UPDATE tokens SET (
    type,
    auth_id,
    visual_number,
    issuer,
    allowed,
    valid,
    whitelist,
    language,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10)
  WHERE uid = $1
  RETURNING id, uid, type, auth_id, visual_number, issuer, allowed, valid, whitelist, language, last_updated
`

type UpdateTokenByUidParams struct {
	Uid          string             `db:"uid" json:"uid"`
	Type         TokenType          `db:"type" json:"type"`
	AuthID       string             `db:"auth_id" json:"authID"`
	VisualNumber sql.NullString     `db:"visual_number" json:"visualNumber"`
	Issuer       string             `db:"issuer" json:"issuer"`
	Allowed      TokenAllowedType   `db:"allowed" json:"allowed"`
	Valid        bool               `db:"valid" json:"valid"`
	Whitelist    TokenWhitelistType `db:"whitelist" json:"whitelist"`
	Language     sql.NullString     `db:"language" json:"language"`
	LastUpdated  time.Time          `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateTokenByUid(ctx context.Context, arg UpdateTokenByUidParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, updateTokenByUid,
		arg.Uid,
		arg.Type,
		arg.AuthID,
		arg.VisualNumber,
		arg.Issuer,
		arg.Allowed,
		arg.Valid,
		arg.Whitelist,
		arg.Language,
		arg.LastUpdated,
	)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Type,
		&i.AuthID,
		&i.VisualNumber,
		&i.Issuer,
		&i.Allowed,
		&i.Valid,
		&i.Whitelist,
		&i.Language,
		&i.LastUpdated,
	)
	return i, err
}
