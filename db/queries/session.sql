-- name: CreateSession :one
INSERT INTO sessions (
    uid,
    start_datetime,
    end_datetime,
    kwh,
    auth_id,
    auth_method,
    location_id,
    meter_id,
    currency,
    total_cost,
    status,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
  RETURNING *;

-- name: GetSessionByUid :one
SELECT * FROM sessions
  WHERE uid = $1;

-- name: UpdateSessionByUid :one
UPDATE sessions SET (
    start_datetime,
    end_datetime,
    kwh,
    auth_id,
    auth_method,
    location_id,
    meter_id,
    currency,
    total_cost,
    status,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
  WHERE uid = $1
  RETURNING *;
