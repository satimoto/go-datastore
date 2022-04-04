-- name: CreateLocation :one
INSERT INTO locations (
    uid, 
    country_code,
    party_id,
    type, 
    name, 
    address, 
    city, 
    postal_code, 
    country, 
    geom, 
    geo_location_id, 
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed, 
    energy_mix_id, 
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
  RETURNING *;

-- name: DeleteLocation :exec
DELETE FROM locations
  WHERE id = $1;

-- name: DeleteLocationByUid :one
DELETE FROM locations
  WHERE uid = $1
  RETURNING *;

-- name: GetLocation :one
SELECT * FROM locations
  WHERE id = $1;

-- name: GetLocationByUid :one
SELECT * FROM locations
  WHERE uid = $1;

-- name: ListLocations :many
SELECT * FROM locations
  ORDER BY name;

-- name: UpdateLocation :one
UPDATE locations SET (
    country_code,
    party_id,
    type, 
    name, 
    address, 
    city, 
    postal_code, 
    country, 
    geom, 
    geo_location_id, 
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
  WHERE id = $1
  RETURNING *;

-- name: UpdateLocationByUid :one
UPDATE locations SET (
    type, 
    name, 
    address, 
    city, 
    postal_code, 
    country, 
    geom, 
    geo_location_id, 
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
  WHERE uid = $1
  RETURNING *;

-- name: UpdateLocationLastUpdated :exec
UPDATE locations SET last_updated = $2
  WHERE id = $1;
