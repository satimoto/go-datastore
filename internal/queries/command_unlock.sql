-- name: CreateCommandUnlock :one
INSERT INTO command_unlocks (
    status,
    location_id,
    evse_uid,
    connector_id,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: GetCommandUnlock :one
SELECT * FROM command_unlocks
  WHERE id = $1;

-- name: UpdateCommandUnlock :one
UPDATE command_unlocks SET (
    status,
    last_updated
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
