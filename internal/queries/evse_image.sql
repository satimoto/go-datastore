-- name: SetEvseImage :exec
INSERT INTO evse_images (evse_id, image_id)
  VALUES ($1, $2);

-- name: DeleteEvseImages :exec
DELETE FROM images im
  USING evse_images ei
  WHERE ei.image_id = im.id AND ei.evse_id = $1;

-- name: ListEvseImages :many
SELECT im.* FROM images im
  INNER JOIN evse_images ei ON ei.image_id = im.id
  WHERE ei.evse_id = $1
  ORDER BY im.id;
