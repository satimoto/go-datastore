-- name: CreateStatusSchedule :one
INSERT INTO status_schedules (
    evse_id, 
    period_begin, 
    period_end, 
    status
  ) VALUES ($1, $2, $3, $3)
  RETURNING *;

-- name: DeleteStatusSchedules :exec
DELETE FROM status_schedules
  WHERE evse_id = $1;

-- name: ListStatusSchedules :many
SELECT * FROM status_schedules
  WHERE evse_id = $1
  ORDER BY id;
