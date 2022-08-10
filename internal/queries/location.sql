-- name: CreateLocation :one
INSERT INTO locations (
    uid, 
    credential_id,
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
    available_evses,
    total_evses,
    is_remote_capable,
    is_rfid_capable,
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed, 
    energy_mix_id, 
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)
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

-- name: GetLocationByLastUpdated :one
SELECT * FROM locations
  WHERE (@credental_id::BIGINT = -1 OR @credental_id::BIGINT = credental_id) AND
    (@country_code::TEXT = '' OR @country_code::TEXT = country_code) AND
    (@party_id::TEXT = '' OR @party_id::TEXT = party_id)
  ORDER BY last_updated DESC
  LIMIT 1;

-- name: GetLocationByUid :one
SELECT * FROM locations
  WHERE uid = $1;

-- name: ListLocations :many
SELECT * FROM locations
  ORDER BY name;

-- name: ListLocationsByGeom :many
SELECT * FROM locations
  WHERE total_evses > 0 AND 
    ST_Intersects(geom, ST_MakeEnvelope(@x_min::FLOAT, @y_min::FLOAT, @x_max::FLOAT, @y_max::FLOAT, 4326)) AND
    (@last_update::TEXT = '' OR last_updated > TO_TIMESTAMP(@last_update::TEXT, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'))
  LIMIT 500;

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
    available_evses,
    total_evses,
    is_remote_capable,
    is_rfid_capable,
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
  WHERE id = $1
  RETURNING *;

-- name: UpdateLocationByUid :one
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
    available_evses,
    total_evses,
    is_remote_capable,
    is_rfid_capable,
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
  WHERE uid = $1
  RETURNING *;

-- name: UpdateLocationAvailability :exec
UPDATE locations SET (
    available_evses, 
    total_evses, 
    is_remote_capable, 
    is_rfid_capable
  ) = ($2, $3, $4, $5)
  WHERE id = $1;

-- name: UpdateLocationLastUpdated :exec
UPDATE locations SET last_updated = $2
  WHERE id = $1;
