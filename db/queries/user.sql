-- name: CreateUser :one
INSERT INTO users (
    linking_key,
    node_key,
    node_address,
    device_token
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: GetUser :one
SELECT * FROM users
  WHERE id = $1;

-- name: GetUserByLinkingKey :one
SELECT * FROM users
  WHERE linking_key = $1;

-- name: GetUserByNodeKey :one
SELECT * FROM users
  WHERE node_key = $1;

-- name: UpdateUser :one
UPDATE users SET (
    node_address,
    device_token
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
