-- name: CreateBusinessDetail :one
INSERT INTO business_details (name, website, logo_id)
  VALUES ($1, $2, $3)
  RETURNING *;

-- name: DeleteBusinessDetail :exec
DELETE FROM business_details
  WHERE id = $1;

-- name: GetBusinessDetail :one
SELECT * FROM business_details
  WHERE id = $1;

-- name: UpdateBusinessDetail :one
UPDATE business_details SET (
    name, 
    website, 
    logo_id
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;
