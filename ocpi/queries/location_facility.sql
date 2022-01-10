-- name: SetLocationFacility :exec
INSERT INTO location_facilities (location_id, facility_id)
  VALUES ($1, $2);

-- name: UnsetLocationFacilities :exec
DELETE FROM location_facilities lf
  WHERE lf.location_id == $1;

-- name: ListLocationFacilities :many
SELECT fa.* FROM facilities fa
  INNER JOIN location_facilities lf ON lf.facility_id == fa.id
  WHERE lf.location_id == $1
  ORDER BY fa.id;
