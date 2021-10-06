-- name: CreateDisplayText :one
INSERT INTO display_texts (language, text)
  VALUES ($1, $2)
  RETURNING *;

-- name: DeleteDisplayText :exec
DELETE FROM display_texts
  WHERE id = $1;

-- name: GetDisplayText :one
SELECT * FROM display_texts
  WHERE id = $1;

-- name: ListDisplayTexts :many
SELECT * FROM display_texts
  ORDER BY id;

-- name: UpdateDisplayText :one
UPDATE display_texts SET (
    language, 
    text
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
