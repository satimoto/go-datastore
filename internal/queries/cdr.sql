-- name: CreateCdr :one
INSERT INTO cdrs (
    uid,
    credential_id,
    country_code, 
    party_id, 
    authorization_id,
    start_date_time,
    stop_date_time,
    auth_id,
    auth_method,
    user_id,
    token_id,
    location_id,
    evse_id,
    connector_id,
    meter_id,
    currency,
    calibration_id,
    total_cost,
    total_energy,
    total_time,
    total_parking_time,
    remark,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
  RETURNING *;

-- name: GetCdrByAuthorizationID :one
SELECT * FROM cdrs
  WHERE authorization_id = @authorization_id::string
  LIMIT 1;

-- name: GetCdrByLastUpdated :one
SELECT * FROM cdrs
  WHERE (@credential_id::BIGINT = -1 OR @credential_id::BIGINT = credential_id) AND
    (@country_code::TEXT = '' OR @country_code::TEXT = country_code) AND
    (@party_id::TEXT = '' OR @party_id::TEXT = party_id)
  ORDER BY last_updated DESC
  LIMIT 1;

-- name: GetCdrByUid :one
SELECT * FROM cdrs
  WHERE uid = $1;

-- name: ListCdrsByCompletedSessionStatus :many
SELECT c.* FROM cdrs c
  INNER JOIN sessions s ON s.authorization_id = c.authorization_id
  INNER JOIN users u ON u.id = s.user_id
  WHERE u.node_id = @node_id::BIGINT AND s.status = 'COMPLETED'
  ORDER BY c.id;

-- name: UpdateCdrIsFlaggedByUid :exec
UPDATE cdrs SET is_flagged = $2
  WHERE uid = $1;
