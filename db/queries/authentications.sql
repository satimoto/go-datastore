-- name: CreateAuthentication :one
INSERT INTO authentications (
    code, 
    action, 
    challenge
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: DeleteAuthentication :exec
DELETE FROM authentications
  WHERE id = $1;

-- name: GetAuthenticationByChallenge :one
SELECT * FROM authentications
  WHERE challenge = $1;

-- name: GetAuthenticationByCode :one
SELECT * FROM authentications
  WHERE code = $1;

-- name: UpdateAuthentication :one
UPDATE authentications SET (
    signature, 
    linking_key
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
