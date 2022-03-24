-- name: CreateEnvironmentalImpact :one
INSERT INTO environmental_impacts (
    energy_mix_id, 
    source, 
    amount
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: DeleteEnvironmentalImpacts :exec
DELETE FROM environmental_impacts
  WHERE energy_mix_id = $1;

-- name: ListEnvironmentalImpacts :many
SELECT * FROM environmental_impacts
  WHERE energy_mix_id = $1
  ORDER BY id;
