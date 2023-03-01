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
    is_intermediate_cdr_capable,
    is_remote_capable,
    is_rfid_capable,
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed, 
    energy_mix_id, 
    is_published,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)
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
  WHERE (@credential_id::BIGINT = -1 OR @credential_id::BIGINT = credential_id) AND
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

-- name: ListLocationsByCountry :many
SELECT * FROM locations
  WHERE is_published = true AND is_removed = false AND total_evses > 0 AND country = $1;

-- name: ListLocationsByGeom :many
SELECT * FROM locations
  WHERE is_published = true AND is_removed = false AND total_evses > 0 AND 
    (@is_experimental::BOOLEAN = true OR is_intermediate_cdr_capable = true) AND 
    (
      (@is_remote_capable::BOOLEAN = true AND is_remote_capable = true) OR 
      (@is_rfid_capable::BOOLEAN = true AND is_rfid_capable = true)
    ) AND
    ST_Intersects(geom, ST_MakeEnvelope(@x_min::FLOAT, @y_min::FLOAT, @x_max::FLOAT, @y_max::FLOAT, 4326)) AND
    (@interval::INT = 0 OR last_updated > NOW() - '1 second'::INTERVAL * @interval::INT)
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
    is_intermediate_cdr_capable,
    is_remote_capable,
    is_rfid_capable,
    is_removed,
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
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
    is_intermediate_cdr_capable,
    is_remote_capable,
    is_rfid_capable,
    is_removed,
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
  WHERE uid = $1
  RETURNING *;

-- name: UpdateLocationLastUpdated :exec
UPDATE locations SET (
    available_evses, 
    total_evses, 
    is_intermediate_cdr_capable,
    is_remote_capable, 
    is_rfid_capable,
    is_removed,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8)
  WHERE id = $1;

-- name: UpdateLocationPublished :exec
UPDATE locations SET is_published = $2
  WHERE id = $1;

-- name: UpdateLocationsPublishedByCredential :exec
UPDATE locations SET is_published = $2
  WHERE credential_id = $1;

-- name: UpdateLocationsPublishedByPartyAndCountryCode :exec
UPDATE locations SET is_published = $3
  WHERE party_id = $1 AND country_code = $2;

-- name: UpdateLocationsRemovedByCredential :exec
UPDATE locations SET is_removed = $2
  WHERE credential_id = $1;

-- name: UpdateLocationsRemovedByPartyAndCountryCode :exec
UPDATE locations SET is_removed = $3
  WHERE party_id = $1 AND country_code = $2;
