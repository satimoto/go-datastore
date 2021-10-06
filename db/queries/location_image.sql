-- name: SetLocationImage :exec
INSERT INTO location_images (location_id, image_id)
  VALUES ($1, $2);

-- name: DeleteLocationImages :exec
DELETE FROM images im
  USING location_images li
  WHERE li.image_id == im.id AND li.location_id == $1;

-- name: ListLocationImages :many
SELECT im.* FROM images im
  INNER JOIN location_images li ON li.image_id == im.id
  WHERE li.location_id == $1
  ORDER BY im.id;
