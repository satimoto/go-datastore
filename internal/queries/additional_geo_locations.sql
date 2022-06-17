-- name: CreateAdditionalGeoLocation :one
INSERT INTO additional_geo_locations (
    location_id,
    display_text_id,
    latitude, 
    latitude_float, 
    longitude, 
    longitude_float 
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: DeleteAdditionalGeoLocations :exec
DELETE FROM additional_geo_locations
  WHERE location_id = $1;

-- name: ListAdditionalGeoLocations :many
SELECT * FROM additional_geo_locations
  WHERE location_id = $1
  ORDER BY id;
