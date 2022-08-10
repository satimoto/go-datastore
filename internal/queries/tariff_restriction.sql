-- name: CreateTariffRestriction :one
INSERT INTO tariff_restrictions (
    start_time,
    end_time,
    start_time_2,
    end_time_2
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: DeleteTariffRestriction :exec
DELETE FROM tariff_restrictions
  WHERE id = $1;

-- name: GetTariffRestriction :one
SELECT * FROM tariff_restrictions
  WHERE id = $1;

-- name: UpdateTariffRestriction :one
UPDATE tariff_restrictions SET (
    start_time,
    end_time,
    start_time_2,
    end_time_2
  ) = ($2, $3, $4, $5)
  WHERE id = $1
  RETURNING *;
