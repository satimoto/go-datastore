-- name: CreateNode :one
INSERT INTO nodes (pubkey, address)
  VALUES ($1, $2)
  RETURNING *;

-- name: DeleteNode :one
DELETE FROM nodes
  WHERE id = $1
  RETURNING *;

-- name: GetNode :one
SELECT * FROM nodes
  WHERE id = $1;

-- name: ListNodes :many
SELECT * FROM nodes
  ORDER BY name;

-- name: UpdateNode :one
UPDATE nodes
  SET pubkey = $2, address = $3
  WHERE id = $1
  RETURNING *;
