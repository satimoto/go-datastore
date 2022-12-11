-- name: CreateCommandStart :one
INSERT INTO command_starts (
    status,
    token_id,
    authorization_id,
    location_id,
    evse_uid,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: GetCommandStart :one
SELECT * FROM command_starts
  WHERE id = $1;

-- name: GetCommandStartByAuthorizationID :one
SELECT * FROM command_starts
  WHERE authorization_id = $1;

-- name: UpdateCommandStart :one
UPDATE command_starts SET (
    status,
    last_updated
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
