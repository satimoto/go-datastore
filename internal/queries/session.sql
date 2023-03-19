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
    user_id,
    token_id,
    location_id,
    evse_id,
    connector_id,
    meter_id,
    currency,
    total_cost,
    status,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
  RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
  WHERE id = $1;

-- name: GetSessionByAuthorizationID :one
SELECT * FROM sessions
  WHERE authorization_id = @authorization_id::TEXT
  LIMIT 1;

-- name: GetSessionByLastUpdated :one
SELECT * FROM sessions
  WHERE (@credential_id::BIGINT = -1 OR @credential_id::BIGINT = credential_id) AND
    (@country_code::TEXT = '' OR @country_code::TEXT = country_code) AND
    (@party_id::TEXT = '' OR @party_id::TEXT = party_id)
  ORDER BY last_updated DESC
  LIMIT 1;

-- name: GetSessionByUid :one
SELECT * FROM sessions
  WHERE uid = $1;

-- name: ListInProgressSessionsByNodeID :many
SELECT s.* FROM sessions s
  INNER JOIN users u ON u.id = s.user_id
  WHERE u.node_id = $1 AND s.status IN ('PENDING', 'ACTIVE')
  ORDER BY s.id;

-- name: ListInProgressSessionsByUserID :many
SELECT * FROM sessions
  WHERE user_id = $1 AND status IN ('PENDING', 'ACTIVE')
  ORDER BY id DESC;

-- name: ListInvoicedSessionsByUserID :many
SELECT * FROM sessions
  WHERE user_id = $1 AND status = 'INVOICED'
  ORDER BY id DESC;

-- name: UpdateSessionByUid :one
UPDATE sessions SET (
    authorization_id,
    start_datetime,
    end_datetime,
    kwh,
    auth_method,
    meter_id,
    currency,
    total_cost,
    status,
    invoice_request_id,
    is_confirmed,
    is_flagged,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
  WHERE uid = $1
  RETURNING *;

-- name: UpdateSessionIsFlaggedByUid :exec
UPDATE sessions SET is_flagged = $2
  WHERE uid = $1;
