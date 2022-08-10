-- name: CreateEvseStatusPeriod :one
INSERT INTO evse_status_periods (
    evse_id, 
    status, 
    start_date, 
    end_date
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: ListEvseStatusPeriods :many
SELECT * FROM evse_status_periods
  WHERE evse_id = $1
  ORDER BY id;
