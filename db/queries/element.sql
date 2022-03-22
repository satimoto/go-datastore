-- name: CreateElement :one
INSERT INTO elements (
    tariff_id,
    restriction_id
  ) VALUES ($1, $2)
  RETURNING *;

-- name: DeleteElements :exec
DELETE FROM elements
  WHERE tariff_id = $1;

-- name: ListElements :many
SELECT * FROM elements
  WHERE tariff_id = $1
  ORDER BY id;
