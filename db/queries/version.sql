-- name: CreateVersion :one
INSERT INTO versions (
    credential_id, 
    version, 
    url
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: DeleteVersions :exec
DELETE FROM versions
  WHERE credential_id = $1;

-- name: GetVersion :one
SELECT * FROM versions
  WHERE id = $1;

-- name: ListVersions :many
SELECT * FROM versions
  WHERE credential_id = $1
  ORDER BY CONCAT(version, REPEAT('.0', 4 - ARRAY_LENGTH(STRING_TO_ARRAY(version, '.'), 1)))::INET DESC;
