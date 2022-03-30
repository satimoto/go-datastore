-- name: CreateCommandUnlock :one
INSERT INTO command_unlocks (
    status,
    location_id,
    evse_uid,
    connector_id
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: UpdateCommandUnlock :one
UPDATE command_unlocks SET status = $2
  WHERE id = $1
  RETURNING *;
