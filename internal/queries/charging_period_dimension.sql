-- name: CreateChargingPeriodDimension :one
INSERT INTO charging_period_dimensions (
    charging_period_id,
    type,
    volume
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: DeleteChargingPeriodDimensions :exec
DELETE FROM charging_period_dimensions cpd
  WHERE cpd.charging_period_id = $1;

-- name: ListChargingPeriodDimensions :many
SELECT cpd.* FROM charging_period_dimensions cpd
  WHERE cpd.charging_period_id = $1
  ORDER BY cpd.id;
