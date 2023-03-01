-- name: SetPoiTag :exec
INSERT INTO poi_tags (poi_id, tag_id)
  VALUES ($1, $2);

-- name: UnsetPoiTags :exec
DELETE FROM poi_tags pt
  WHERE pt.poi_id = $1;

-- name: ListPoiTags :many
SELECT t.* FROM tags t
  INNER JOIN poi_tags pt ON pt.tag_id = t.id
  WHERE pt.poi_id = $1
  ORDER BY t.id;
