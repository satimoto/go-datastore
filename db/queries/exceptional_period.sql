-- name: CreateExceptionalPeriod :one
INSERT INTO exceptional_periods (
    opening_time_id, 
    period_type, 
    period_begin, 
    period_end
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: DeleteExceptionalOpeningPeriods :exec
DELETE FROM exceptional_periods
  WHERE opening_time_id = $1 AND period_type = 'OPENING';

-- name: DeleteExceptionalClosingPeriods :exec
DELETE FROM exceptional_periods
  WHERE opening_time_id = $1 AND period_type = 'CLOSING';

-- name: ListExceptionalOpeningPeriods :many
SELECT * FROM exceptional_periods
  WHERE opening_time_id = $1 AND period_type = 'OPENING'
  ORDER BY id;

-- name: ListExceptionalClosingPeriods :many
SELECT * FROM exceptional_periods
  WHERE opening_time_id = $1 AND period_type = 'CLOSING'
  ORDER BY id;
