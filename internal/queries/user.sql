-- name: CreateUser :one
INSERT INTO users (
    commission_percent,
    device_token,
    linking_pubkey,
    node_id,
    pubkey,
    referral_code,
    circuit_user_id
  ) VALUES ($1, $2, $3, $4, $5, $6, $7)
  RETURNING *;

-- name: GetUser :one
SELECT * FROM users
  WHERE id = $1;

-- name: GetUserByDeviceToken :one
SELECT * FROM users
  WHERE device_token = $1;

-- name: GetUserByLinkingPubkey :one
SELECT * FROM users
  WHERE linking_pubkey = $1;

-- name: GetUserByPubkey :one
SELECT * FROM users
  WHERE pubkey = $1;

-- name: GetUserByReferralCode :one
SELECT * FROM users
  WHERE referral_code = $1;

-- name: GetUserBySessionID :one
SELECT u.* FROM users u
  INNER JOIN sessions s ON s.user_id = u.id
  WHERE s.id = $1;

-- name: GetUserByTokenID :one
SELECT u.* FROM users u
  INNER JOIN tokens t ON t.user_id = u.id
  WHERE t.id = $1;

-- name: UpdateUser :one
UPDATE users SET (
    commission_percent,
    device_token,
    linking_pubkey,
    node_id,
    pubkey,
    is_restricted,
    referral_code,
    circuit_user_id,
    name,
    address,
    postal_code,
    city
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
  WHERE id = $1
  RETURNING *;

-- name: UpdateUserByPubkey :one
UPDATE users SET 
  last_active_date = $2
  WHERE pubkey = $1
  RETURNING *;
