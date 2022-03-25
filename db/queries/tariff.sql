-- name: CreateTariff :one
INSERT INTO tariffs (
    uid, 
    currency, 
    tariff_alt_url, 
    energy_mix_id, 
    last_updated
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: DeleteTariffByUid :exec
DELETE FROM tariffs
  WHERE uid = $1;

-- name: GetTariffByUid :one
SELECT * FROM tariffs
  WHERE uid = $1;

-- name: UpdateTariffByUid :one
UPDATE tariffs SET (
    currency, 
    tariff_alt_url,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5)
  WHERE uid = $1
  RETURNING *;
