-- name: CreateTariff :one
INSERT INTO tariffs (
    uid, 
    credential_id,
    country_code,
    party_id,
    cdr_id,
    currency, 
    tariff_alt_url, 
    energy_mix_id, 
    tariff_restriction_id,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
  RETURNING *;

-- name: DeleteTariffByUid :exec
DELETE FROM tariffs
  WHERE uid = $1 AND cdr_id IS NULL;

-- name: GetTariff :one
SELECT * FROM tariffs
  WHERE id = $1;

-- name: GetTariffByLastUpdated :one
SELECT * FROM tariffs
  WHERE (@credental_id::BIGINT = -1 OR @credental_id::BIGINT = credental_id) AND
    (@country_code::TEXT = '' OR @country_code::TEXT = country_code) AND
    (@party_id::TEXT = '' OR @party_id::TEXT = party_id)
  ORDER BY last_updated DESC
  LIMIT 1;

-- name: GetTariffByUid :one
SELECT * FROM tariffs
  WHERE uid = $1 AND cdr_id IS NULL;

-- name: ListTariffsByCdr :many
SELECT * FROM tariffs
  WHERE cdr_id = $1
  ORDER BY id;

-- name: UpdateTariffByUid :one
UPDATE tariffs SET (
    country_code,
    party_id,
    currency, 
    tariff_alt_url,
    energy_mix_id, 
    tariff_restriction_id,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8)
  WHERE uid = $1 AND cdr_id IS NULL
  RETURNING *;
