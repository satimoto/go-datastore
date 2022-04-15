-- name: CreateSession :one
INSERT INTO sessions (
    uid,
    credential_id,
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
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
  RETURNING *;

-- name: GetSessionByLastUpdated :one
SELECT * FROM sessions
  WHERE (@credental_id::BIGINT = -1 OR @credental_id::BIGINT = credental_id) AND
    (@country_code::TEXT = '' OR @country_code::TEXT = country_code) AND
    (@party_id::TEXT = '' OR @party_id::TEXT = party_id)
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
