-- name: CreatePriceComponent :one
INSERT INTO price_components (
    element_id,
    type,
    price,
    step_size
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: DeletePriceComponents :exec
DELETE FROM price_components pc
  USING elements e
  WHERE pc.element_id == e.id AND e.tariff_id == $1;

-- name: ListPriceComponents :many
SELECT * FROM price_components
  WHERE element_id = $1
  ORDER BY id;
