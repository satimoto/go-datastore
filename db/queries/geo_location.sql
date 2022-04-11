-- name: CreateGeoLocation :one
INSERT INTO geo_locations (
     latitude, 
     latitude_float, 
     longitude, 
     longitude_float, 
     name
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: DeleteGeoLocation :exec
DELETE FROM geo_locations
  WHERE id = $1;

-- name: GetGeoLocation :one
SELECT * FROM geo_locations
  WHERE id = $1;

-- name: ListGeoLocations :many
SELECT * FROM geo_locations
  ORDER BY name;

-- name: UpdateGeoLocation :one
UPDATE geo_locations SET (
    latitude, 
    latitude_float, 
    longitude, 
    longitude_float, 
    name
  ) = ($2, $3, $4, $5, $6)
  WHERE id = $1
  RETURNING *;
