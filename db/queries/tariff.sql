-- name: CreateTariff :one
INSERT INTO tariffs (
    uid, 
    country_code,
    party_id,
    cdr_id,
    currency, 
    tariff_alt_url, 
    energy_mix_id, 
    tariff_restriction_id,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  RETURNING *;

-- name: DeleteTariffByUid :exec
DELETE FROM tariffs
  WHERE uid = $1 AND cdr_id IS NULL;

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
