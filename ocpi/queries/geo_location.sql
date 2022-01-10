-- name: CreateGeoLocation :one
INSERT INTO geo_locations (latitude, longitude, name)
  VALUES ($1, $2, $3)
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
    longitude, 
    name
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;
