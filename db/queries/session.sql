-- name: CreateSession :one
INSERT INTO sessions (
    uid,
    country_code,
    party_id,
    authorization_id,
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
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
  RETURNING *;

-- name: GetSessionByIdentityOrderByLastUpdated :one
SELECT * FROM sessions
  WHERE country_code = $1 AND party_id = $2
  ORDER BY last_updated DESC
  LIMIT 1;

-- name: GetSessionByUid :one
SELECT * FROM sessions
  WHERE uid = $1;

-- name: UpdateSessionByUid :one
UPDATE sessions SET (
    country_code,
    party_id,
    authorization_id,
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
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
  WHERE uid = $1
  RETURNING *;
