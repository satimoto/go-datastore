-- name: GetFacility :one
SELECT * FROM facilities
  WHERE id = $1;

-- name: ListFacilities :many
SELECT * FROM facilities
  ORDER BY id;
