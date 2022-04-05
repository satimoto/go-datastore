-- name: CreateElementRestriction :one
INSERT INTO element_restrictions (
    start_time,
    end_time,
    start_date,
    end_date,
    min_kwh,
    max_kwh,
    min_power,
    max_power,
    min_duration,
    max_duration
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
  RETURNING *;

-- name: DeleteElementRestrictions :exec
DELETE FROM element_restrictions r
  USING elements e
  WHERE r.id == e.restriction_id AND e.tariff_id == $1;

-- name: GetElementRestriction :one
SELECT * FROM element_restrictions
  WHERE id = $1;

-- name: UpdateElementRestriction :one
UPDATE element_restrictions SET (
    start_time,
    end_time,
    start_date,
    end_date,
    min_kwh,
    max_kwh,
    min_power,
    max_power,
    min_duration,
    max_duration
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE id = $1
  RETURNING *;
