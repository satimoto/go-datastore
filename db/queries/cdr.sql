-- name: CreateCdr :one
INSERT INTO cdrs (
    uid,
    country_code, 
    party_id, 
    authorization_id,
    start_date_time,
    stop_date_time,
    auth_id,
    auth_method,
    location_id,
    meter_id,
    currency,
    calibration_id,
    total_cost,
    total_energy,
    total_time,
    total_parking_time,
    remark,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
  RETURNING *;

-- name: GetCdrByUid :one
SELECT * FROM cdrs
  WHERE uid = $1;
