-- name: GetParkingRestriction :one
SELECT * FROM parking_restrictions
  WHERE id = $1;

-- name: ListParkingRestrictions :many
SELECT * FROM parking_restrictions
  ORDER BY id;
