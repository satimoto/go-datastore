-- name: CreatePriceComponent :one
INSERT INTO price_components (
    element_id,
    type,
    currency,
    price,
    step_size,
    price_rounding_id,
    step_rounding_id,
    exact_price_component
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING *;

-- name: DeletePriceComponents :exec
DELETE FROM price_components pc
  USING elements e
  WHERE pc.element_id = e.id AND e.tariff_id = $1;

-- name: ListPriceComponents :many
SELECT * FROM price_components
  WHERE element_id = $1
  ORDER BY id;
