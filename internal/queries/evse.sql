-- name: CreateEvse :one
INSERT INTO evses (
    location_id, 
    uid, 
    evse_id, 
    status, 
    floor_level, 
    geom, 
    geo_location_id, 
    is_remote_capable,
    is_rfid_capable,
    physical_reference, 
    last_updated)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  RETURNING *;

-- name: DeleteEvse :exec
DELETE FROM evses
  WHERE id = $1;

-- name: DeleteEvseByUid :exec
DELETE FROM evses
  WHERE uid = $1;

-- name: GetEvse :one
SELECT * FROM evses
  WHERE id = $1;

-- name: GetEvseByEvseId :one
SELECT * FROM evses
  WHERE evse_id = $1;

-- name: GetEvseByUid :one
SELECT * FROM evses
  WHERE uid = $1;

-- name: ListEvses :many
SELECT * FROM evses
  WHERE location_id = $1
  ORDER BY id;

-- name: UpdateEvse :one
UPDATE evses SET (
    evse_id, 
    status, 
    floor_level, 
    geom, 
    geo_location_id, 
    is_remote_capable,
    is_rfid_capable,
    physical_reference, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10)
  WHERE id = $1
  RETURNING *;

-- name: UpdateEvseByUid :one
UPDATE evses SET (
    evse_id, 
    status, 
    floor_level, 
    geom, 
    geo_location_id, 
    is_remote_capable,
    is_rfid_capable,
    physical_reference, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10)
  WHERE uid = $1
  RETURNING *;

-- name: UpdateEvseLastUpdated :exec
UPDATE evses SET last_updated = $2
  WHERE id = $1;
