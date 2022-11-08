-- name: CreateTokenAuthorization :one
INSERT INTO token_authorizations (
    token_id,
    authorization_id,
    authorized,
    signing_key,
    country_code,
    party_id,
    location_id
  ) VALUES ($1, $2, $3, $4, $5, $6, $7)
  RETURNING *;

-- name: GetTokenAuthorizationByAuthorizationID :one
SELECT * FROM token_authorizations
  WHERE authorization_id = $1;

-- name: GetLastTokenAuthorizationByTokenID :one
SELECT * FROM token_authorizations
  WHERE token_id = $1 
  ORDER BY id DESC
  LIMIT 1;

-- name: UpdateTokenAuthorizationByAuthorizationID :one
UPDATE token_authorizations SET (
    authorized,
    country_code,
    party_id
  ) = ($2, $3, $4)
  WHERE authorization_id = $1
  RETURNING *;
