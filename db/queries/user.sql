-- name: CreateUser :one
INSERT INTO users (device_token, node_id)
  VALUES ($1, $2)
  RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
  WHERE id = $1
  RETURNING *;

-- name: GetUser :one
SELECT * FROM users
  WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
  ORDER BY id;

-- name: UpdateUser :one
UPDATE users
  SET device_token = $2
  WHERE id = $1
  RETURNING *;
