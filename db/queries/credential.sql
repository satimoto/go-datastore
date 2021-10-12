-- name: CreateCredential :one
INSERT INTO credentials (
    client_token, 
    server_token, 
    url, 
    business_detail_id, 
    party_id, 
    country_code, 
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7)
  RETURNING *;

-- name: DeleteCredential :exec
DELETE FROM credentials
  WHERE id = $1;

-- name: GetCredential :one
SELECT * FROM credentials
  WHERE id = $1;

-- name: GetCredentialByPartyAndCountryCode :one
SELECT * FROM credentials
  WHERE party_id = $1 AND country_code = $2;

-- name: UpdateCredential :one
UPDATE credentials SET (
    client_token, 
    server_token, 
    url, 
    party_id, 
    country_code,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7)
  WHERE id = $1
  RETURNING *;
