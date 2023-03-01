-- name: CreateTag :one
INSERT INTO tags (
    key, 
    value
  ) VALUES ($1, $2)
  RETURNING *;

-- name: GetTagByKeyValue :one
SELECT * FROM tags
  WHERE key = $1 AND value = $2
  LIMIT 1;
