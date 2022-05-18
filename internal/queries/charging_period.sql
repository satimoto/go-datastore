-- name: CreateChargingPeriod :one
INSERT INTO charging_periods (
    start_date_time
  ) VALUES ($1)
  RETURNING *;
