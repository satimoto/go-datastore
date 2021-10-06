-- name: SetEvseDirection :exec
INSERT INTO evse_directions (evse_id, display_text_id)
  VALUES ($1, $2);

-- name: DeleteEvseDirections :exec
DELETE FROM display_texts dt
  USING evse_directions ed
  WHERE ed.display_text_id == dt.id AND ed.evse_id == $1;

-- name: ListEvseDirections :many
SELECT dt.* FROM display_texts dt
  INNER JOIN evse_directions ed ON ed.display_text_id == dt.id
  WHERE ed.evse_id == $1
  ORDER BY dt.id;
