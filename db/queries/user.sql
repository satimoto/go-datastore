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

-- name: GetUserByPubkey :one
SELECT * FROM users
  WHERE pubkey = $1;

-- name: UpdateUser :one
UPDATE users SET (
    device_token,
    node_id
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
