-- name: CreatePriceComponentRounding :one
INSERT INTO price_component_roundings (
    granularity,
    rule
  ) VALUES ($1, $2)
  RETURNING *;

-- name: DeletePriceComponentRoundings :exec
DELETE FROM price_component_roundings pcr
  USING elements e, price_components pc
  WHERE e.tariff_id = $1 AND e.id = pc.element_id AND 
    (pc.price_rounding_id = pcr.id OR pc.step_rounding_id = pcr.id);

-- name: GetPriceComponentRounding :one
SELECT * FROM price_component_roundings
  WHERE id = $1;
