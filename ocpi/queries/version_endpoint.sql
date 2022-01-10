-- name: CreateVersionEndpoint :one
INSERT INTO version_endpoints (
    version_id, 
    identifier, 
    url
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: DeleteVersionEndpoints :exec
DELETE FROM version_endpoints
  WHERE version_id = $1;

-- name: GetVersionEndpoint :one
SELECT * FROM version_endpoints
  WHERE id = $1;

-- name: ListVersionEndpoints :many
SELECT * FROM version_endpoints
  WHERE version_id = $1;
