-- name: CreateEnergyMix :one
INSERT INTO energy_mixes (
    is_green_energy, 
    supplier_name, 
    energy_product_name
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: DeleteEnergyMix :exec
DELETE FROM energy_mixes
  WHERE id = $1;

-- name: GetEnergyMix :one
SELECT * FROM energy_mixes
  WHERE id = $1;

-- name: UpdateEnergyMix :one
UPDATE energy_mixes SET (
    is_green_energy, 
    supplier_name, 
    energy_product_name
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;
