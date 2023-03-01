-- name: CreateSessionUpdate :one
INSERT INTO session_updates (
    session_id,
    user_id,
    kwh,
    total_cost,
    status,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: ListSessionUpdatesBySessionID :many
SELECT * FROM session_updates
  WHERE session_id = $1
  ORDER BY id;
