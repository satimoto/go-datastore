-- name: CreateCdr :one
INSERT INTO cdrs (
    uid,
    start_date_time,
    stop_date_time,
    auth_id,
    auth_method,
    location_id,
    meter_id,
    currency,
    total_cost,
    total_energy,
    total_time,
    total_parking_time,
    remark,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
  RETURNING *;

-- name: GetCdrByUid :one
SELECT * FROM cdrs
  WHERE uid = $1;
