-- name: CreateCommandStart :one
INSERT INTO command_starts (
    status,
    token_id,
    location_id,
    evse_uid
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: UpdateCommandStart :one
UPDATE command_starts SET status = $2
  WHERE id = $1
  RETURNING *;
