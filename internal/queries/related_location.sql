-- name: SetRelatedLocation :exec
INSERT INTO related_locations (location_id, geo_location_id)
  VALUES ($1, $2);

-- name: DeleteRelatedLocations :exec
DELETE FROM geo_locations gl
  USING related_locations rl
  WHERE rl.geo_location_id = gl.id AND rl.location_id = $1;

-- name: ListRelatedLocations :many
SELECT gl.* FROM geo_locations gl
  INNER JOIN related_locations rl ON rl.geo_location_id = gl.id
  WHERE rl.location_id = $1
  ORDER BY gl.id;
