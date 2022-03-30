-- name: CreateTariff :one
INSERT INTO tariffs (
    uid, 
    cdr_id,
    currency, 
    tariff_alt_url, 
    energy_mix_id, 
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6)
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
    currency, 
    tariff_alt_url,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5)
  WHERE uid = $1 AND cdr_id IS NULL
  RETURNING *;
