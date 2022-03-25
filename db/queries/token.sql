-- name: CreateToken :one
INSERT INTO tokens (
    uid, 
    type,
    auth_id,
    visual_number,
    issuer,
    valid,
    whitelist,
    language,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  RETURNING *;

-- name: DeleteTokenByUid :exec
DELETE FROM tokens
  WHERE uid = $1;

-- name: GetTokenByUid :one
SELECT * FROM tokens
  WHERE uid = $1;

-- name: UpdateTokenByUid :one
UPDATE tokens SET (
    type,
    auth_id,
    visual_number,
    issuer,
    valid,
    whitelist,
    language,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9)
  WHERE uid = $1
  RETURNING *;
