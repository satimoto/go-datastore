-- name: CreateCommandStop :one
INSERT INTO command_stops (
    status,
    session_id
  ) VALUES ($1, $2)
  RETURNING *;

-- name: UpdateCommandStop :one
UPDATE command_stops SET status = $2
  WHERE id = $1
  RETURNING *;
