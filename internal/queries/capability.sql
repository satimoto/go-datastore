-- name: GetCapability :one
SELECT * FROM capabilities
  WHERE id = $1;

-- name: ListCapabilities :many
SELECT * FROM capabilities
  ORDER BY id;
