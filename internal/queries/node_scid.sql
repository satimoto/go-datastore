-- name: CountNodeScids :one
SELECT COUNT(*) FROM node_scids
  WHERE node_id = $1;

-- name: CreateNodeScid :one
INSERT INTO node_scids (
    node_id, 
    scid
  ) VALUES ($1, $2)
  RETURNING *;

-- name: DeleteNodeScid :exec
DELETE FROM node_scids
  WHERE id = $1;

-- name: GetNodeScid :one
SELECT * FROM node_scids
  WHERE node_id = $1
  ORDER BY id ASC
  LIMIT 1;
