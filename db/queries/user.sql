-- name: CreateUser :one
INSERT INTO users (
    device_token,
    linking_pubkey,
    node_id,
    pubkey
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: GetUser :one
SELECT * FROM users
  WHERE id = $1;

-- name: GetUserByDeviceToken :one
SELECT * FROM users
  WHERE device_token = $1;

-- name: GetUserByLinkingPubkey :one
SELECT * FROM users
  WHERE linking_pubkey = $1;

-- name: GetUserByTokenID :one
SELECT u.* FROM users u
  INNER JOIN tokens t ON t.user_id = u.id
  WHERE t.id = $1;

-- name: GetUserByPubkey :one
SELECT * FROM users
  WHERE pubkey = $1;

-- name: UpdateUser :one
UPDATE users SET (
    device_token,
    linking_pubkey,
    node_id,
    pubkey
  ) = ($2, $3, $4, $5)
  WHERE id = $1
  RETURNING *;
