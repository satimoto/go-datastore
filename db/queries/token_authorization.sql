-- name: CreateTokenAuthorization :one
INSERT INTO token_authorizations (
    token_id,
    authorization_id,
    country_code,
    party_id,
    location_id
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: GetTokenAuthorizationByAuthorizationID :one
SELECT * FROM token_authorizations
  WHERE authorization_id = $1;

-- name: UpdateTokenAuthorizationByAuthorizationID :one
UPDATE token_authorizations SET (
    country_code,
    party_id
  ) = ($2, $3)
  WHERE authorization_id = $1
  RETURNING *;
