-- name: CreateCommandStop :one
INSERT INTO command_stops (
    status,
    session_id,
    last_updated
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: GetCommandStop :one
SELECT * FROM command_stops
  WHERE id = $1;

-- name: GetCommandStopBySessionID :one
SELECT * FROM command_stops
  WHERE session_id = $1 AND status = 'REQUESTED';

-- name: UpdateCommandStop :one
UPDATE command_stops SET (
    status,
    last_updated
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
