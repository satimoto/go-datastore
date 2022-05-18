-- name: SetLocationDirection :exec
INSERT INTO location_directions (location_id, display_text_id)
  VALUES ($1, $2);

-- name: DeleteLocationDirections :exec
DELETE FROM display_texts dt
  USING location_directions ld
  WHERE ld.display_text_id = dt.id AND ld.location_id = $1;

-- name: ListLocationDirections :many
SELECT dt.* FROM display_texts dt
  INNER JOIN location_directions ld ON ld.display_text_id = dt.id
  WHERE ld.location_id = $1
  ORDER BY dt.id;
