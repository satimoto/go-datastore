-- name: CreateTokenAuthorization :one
INSERT INTO token_authorizations (
    token_id,
    authorization_id,
    location_id
  ) VALUES ($1, $2, $3)
  RETURNING *;
