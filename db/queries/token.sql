-- name: CreateToken :one
INSERT INTO tokens (
    uid, 
    user_id,
    type,
    auth_id,
    visual_number,
    issuer,
    allowed,
    valid,
    whitelist,
    language,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  RETURNING *;

-- name: DeleteTokenByUid :exec
DELETE FROM tokens
  WHERE uid = $1;

-- name: GetToken :one
SELECT * FROM tokens
  WHERE id = $1;

-- name: GetTokenByAuthID :one
SELECT * FROM tokens
  WHERE auth_id = $1;

-- name: GetTokenByUid :one
SELECT * FROM tokens
  WHERE uid = $1;

-- name: GetTokenByUserID :one
SELECT * FROM tokens
  WHERE user_id = $1 AND type = $2
  LIMIT 1;

-- name: ListTokens :many
SELECT * FROM tokens
  WHERE 
    (@filter_date_from::text = '' or last_updated > @filter_date_from::text) and 
    (@filter_date_to::text = '' or last_updated < @filter_date_to::text)
  ORDER BY id
  LIMIT @filter_limit::bigint
  OFFSET @filter_offset::bigint;

-- name: ListTokensByUserID :many
SELECT * FROM tokens
  WHERE user_id = $1;

-- name: UpdateTokenByUid :one
UPDATE tokens SET (
    type,
    auth_id,
    visual_number,
    issuer,
    allowed,
    valid,
    whitelist,
    language,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10)
  WHERE uid = $1
  RETURNING *;
