// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    device_token,
    linking_pubkey,
    node_id,
    pubkey
  ) VALUES ($1, $2, $3, $4)
  RETURNING id, linking_pubkey, pubkey, device_token, node_id
`

type CreateUserParams struct {
	DeviceToken   string        `db:"device_token" json:"deviceToken"`
	LinkingPubkey string        `db:"linking_pubkey" json:"linkingPubkey"`
	NodeID        sql.NullInt64 `db:"node_id" json:"nodeID"`
	Pubkey        string        `db:"pubkey" json:"pubkey"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.DeviceToken,
		arg.LinkingPubkey,
		arg.NodeID,
		arg.Pubkey,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.LinkingPubkey,
		&i.Pubkey,
		&i.DeviceToken,
		&i.NodeID,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, linking_pubkey, pubkey, device_token, node_id FROM users
  WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.LinkingPubkey,
		&i.Pubkey,
		&i.DeviceToken,
		&i.NodeID,
	)
	return i, err
}

const getUserByLinkingPubkey = `-- name: GetUserByLinkingPubkey :one
SELECT id, linking_pubkey, pubkey, device_token, node_id FROM users
  WHERE linking_pubkey = $1
`

func (q *Queries) GetUserByLinkingPubkey(ctx context.Context, linkingPubkey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByLinkingPubkey, linkingPubkey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.LinkingPubkey,
		&i.Pubkey,
		&i.DeviceToken,
		&i.NodeID,
	)
	return i, err
}

const getUserByPubkey = `-- name: GetUserByPubkey :one
SELECT id, linking_pubkey, pubkey, device_token, node_id FROM users
  WHERE pubkey = $1
`

func (q *Queries) GetUserByPubkey(ctx context.Context, pubkey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByPubkey, pubkey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.LinkingPubkey,
		&i.Pubkey,
		&i.DeviceToken,
		&i.NodeID,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET (
    node_id, 
    device_token
  ) = ($2, $3)
  WHERE id = $1
  RETURNING id, linking_pubkey, pubkey, device_token, node_id
`

type UpdateUserParams struct {
	ID          int64         `db:"id" json:"id"`
	NodeID      sql.NullInt64 `db:"node_id" json:"nodeID"`
	DeviceToken string        `db:"device_token" json:"deviceToken"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.NodeID, arg.DeviceToken)
	var i User
	err := row.Scan(
		&i.ID,
		&i.LinkingPubkey,
		&i.Pubkey,
		&i.DeviceToken,
		&i.NodeID,
	)
	return i, err
}
