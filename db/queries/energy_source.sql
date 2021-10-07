-- name: CreateEnergySource :one
INSERT INTO energy_sources (
    energy_mix_id, 
    source, 
    percentage
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: DeleteEnergySources :exec
DELETE FROM energy_sources
  WHERE energy_mix_id = $1;

-- name: ListEnergySources :many
SELECT * FROM energy_sources
  WHERE energy_mix_id = $1
  ORDER BY id;
