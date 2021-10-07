-- name: CreateImage :one
INSERT INTO images (
    url, 
    thumbnail, 
    category, 
    type, 
    width, 
    height
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: DeleteImage :exec
DELETE FROM images
  WHERE id = $1;

-- name: GetImage :one
SELECT * FROM images
  WHERE id = $1;

-- name: UpdateImage :one
UPDATE images SET (
    url, 
    thumbnail, 
    category, 
    type, 
    width, 
    height
  ) = ($2, $3, $4, $5, $6, $7)
  WHERE id = $1
  RETURNING *;
