-- name: CreateParty :one
INSERT INTO parties (
    credential_id,
    country_code,
    party_id,
    is_intermediate_cdr_capable,
    publish_location,
    publish_null_tariff
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: GetParty :one
SELECT * FROM parties
  WHERE id = $1;

-- name: GetPartyByCredential :one
SELECT * FROM parties
  WHERE credential_id = $1 AND country_code = $2 AND party_id = $3;

-- name: UpdateParty :one
UPDATE parties SET (
    is_intermediate_cdr_capable,
    publish_location,
    publish_null_tariff
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;

-- name: UpdatePartyByCredential :one
UPDATE parties SET (
    is_intermediate_cdr_capable,
    publish_location,
    publish_null_tariff
  ) = ($4, $5, $6)
  WHERE credential_id = $1 AND country_code = $2 AND party_id = $3
  RETURNING *;
