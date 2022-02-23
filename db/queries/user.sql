-- name: CreateUser :one
INSERT INTO users (
    linking_pubkey,
    node_pubkey,
    device_token
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: GetUser :one
SELECT * FROM users
  WHERE id = $1;

-- name: GetUserByLinkingPubkey :one
SELECT * FROM users
  WHERE linking_pubkey = $1;

-- name: GetUserByNodePubkey :one
SELECT * FROM users
  WHERE node_pubkey = $1;

-- name: UpdateUser :one
UPDATE users SET device_token = $2
  WHERE id = $1
  RETURNING *;
