-- name: CreateOpeningTime :one
INSERT INTO opening_times (
    twentyfourseven
  ) VALUES ($1)
  RETURNING *;

-- name: DeleteOpeningTime :exec
DELETE FROM opening_times
  WHERE id = $1;

-- name: GetOpeningTime :one
SELECT * FROM opening_times
  WHERE id = $1;

-- name: UpdateOpeningTime :one
UPDATE opening_times 
  SET twentyfourseven = $2
  WHERE id = $1
  RETURNING *;
