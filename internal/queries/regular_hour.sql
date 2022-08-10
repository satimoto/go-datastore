-- name: CreateRegularHour :one
INSERT INTO regular_hours (
    opening_time_id, 
    weekday, 
    period_begin, 
    period_end
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: DeleteRegularHours :exec
DELETE FROM regular_hours
  WHERE opening_time_id = $1;

-- name: ListRegularHours :many
SELECT * FROM regular_hours
  WHERE opening_time_id = $1
  ORDER BY id;
