-- name: CreateCommandStart :one
INSERT INTO command_starts (
    status,
    token_id,
    authorization_id,
    location_id,
    evse_uid
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: GetCommandStart :one
SELECT * FROM command_starts
  WHERE id = $1;

-- name: UpdateCommandStart :one
UPDATE command_starts SET status = $2
  WHERE id = $1
  RETURNING *;
